// Statping
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/statping/statping
//
// The licenses for most software and other practical works are designed
// to take away your freedom to share and change the works.  By contrast,
// the GNU General Public License is intended to guarantee your freedom to
// share and change all versions of a program--to make sure it remains free
// software for all its users.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"flag"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/statping/statping/types/core"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/statping/statping/source"

	"github.com/pkg/errors"
	"github.com/statping/statping/handlers"
	"github.com/statping/statping/types/configs"
	"github.com/statping/statping/types/services"
	"github.com/statping/statping/utils"
)

var (
	// VERSION stores the current version of Statping
	VERSION string
	// COMMIT stores the git commit hash for this version of Statping
	COMMIT      string
	ipAddress   string
	envFile     string
	verboseMode int
	port        int
	log         = utils.Log.WithField("type", "cmd")
	httpServer  = make(chan bool)

	confgs *configs.DbConfig
)

// parseFlags will parse the application flags
// -ip = 0.0.0.0 IP address for outgoing HTTP server
// -port = 8080 Port number for outgoing HTTP server
// environment variables WILL overwrite flags
func parseFlags() {
	envPort := utils.Getenv("PORT", 8080).(int)
	envIpAddress := utils.Getenv("IP", "0.0.0.0").(string)
	envVerbose := utils.Getenv("VERBOSE", 2).(int)

	flag.StringVar(&ipAddress, "ip", envIpAddress, "IP address to run the Statping HTTP server")
	flag.StringVar(&envFile, "env", "", "IP address to run the Statping HTTP server")
	flag.IntVar(&port, "port", envPort, "Port to run the HTTP server")
	flag.IntVar(&verboseMode, "verbose", envVerbose, "Run in verbose mode to see detailed logs (1 - 4)")
	flag.Parse()
}

func exit(err error) {
	sentry.CaptureException(err)
	log.Fatalln(err)
	Close()
	os.Exit(2)
}

func init() {
	core.New(VERSION)
}

// main will run the Statping application
func main() {
	var err error
	go sigterm()

	parseFlags()

	if err := source.Assets(); err != nil {
		exit(err)
	}

	utils.VerboseMode = verboseMode

	if err := utils.InitLogs(); err != nil {
		log.Errorf("Statping Log Error: %v\n", err)
	}

	args := flag.Args()

	if len(args) >= 1 {
		err := catchCLI(args)
		if err != nil {
			if err.Error() == "end" {
				os.Exit(0)
				return
			}
			exit(err)
		}
	}
	log.Info(fmt.Sprintf("Starting Statping v%v", VERSION))

	if err := updateDisplay(); err != nil {
		log.Warnln(err)
	}

	errorEnv := utils.Getenv("GO_ENV", "production").(string)
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:         errorReporter,
		Environment: errorEnv,
	}); err != nil {
		log.Errorln(err)
	}

	confgs, err = configs.LoadConfigs()
	if err != nil {
		if err := SetupMode(); err != nil {
			exit(err)
		}
	}

	if err = configs.ConnectConfigs(confgs); err != nil {
		exit(err)
	}

	if !confgs.Db.HasTable("core") {
		var srvs int64
		if confgs.Db.HasTable(&services.Service{}) {
			confgs.Db.Model(&services.Service{}).Count(&srvs)
			if srvs > 0 {
				exit(errors.Wrap(err, "there are already services setup."))
				return
			}
		}

		if err := confgs.DropDatabase(); err != nil {
			exit(errors.Wrap(err, "error dropping database"))
		}

		if err := confgs.CreateDatabase(); err != nil {
			exit(errors.Wrap(err, "error creating database"))
		}

		if err := configs.CreateAdminUser(confgs); err != nil {
			exit(errors.Wrap(err, "error creating default admin user"))
		}

		if err := configs.TriggerSamples(); err != nil {
			exit(errors.Wrap(err, "error creating database"))
		}

	}

	if err = confgs.DatabaseChanges(); err != nil {
		exit(err)
	}

	if err := confgs.MigrateDatabase(); err != nil {
		exit(err)
	}

	//log.Infoln("Migrating Notifiers...")
	//if err := notifier.Migrate(); err != nil {
	//	exit(errors.Wrap(err, "error migrating notifiers"))
	//}
	//log.Infoln("Notifiers Migrated")

	if err := mainProcess(); err != nil {
		exit(err)
	}
}

// Close will gracefully stop the database connection, and log file
func Close() {
	sentry.Flush(3 * time.Second)
	utils.CloseLogs()
	confgs.Close()
}

func SetupMode() error {
	return handlers.RunHTTPServer(ipAddress, port)
}

// sigterm will attempt to close the database connections gracefully
func sigterm() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	fmt.Println("Shutting down Statping")
	Close()
	os.Exit(0)
}

// mainProcess will initialize the Statping application and run the HTTP server
func mainProcess() error {
	if err := services.ServicesFromEnvFile(); err != nil {
		errStr := "error 'SERVICE' environment variable"
		log.Errorln(errStr)
		return errors.Wrap(err, errStr)
	}

	if err := InitApp(); err != nil {
		return err
	}

	if err := handlers.RunHTTPServer(ipAddress, port); err != nil {
		log.Fatalln(err)
		return errors.Wrap(err, "http server")
	}
	return nil
}

func StartHTTPServer() {
	httpServer = make(chan bool)
	go httpServerProcess(httpServer)
}

func StopHTTPServer() {

}

func httpServerProcess(process <-chan bool) {
	for {
		select {
		case <-process:
			fmt.Println("HTTP Server has stopped")
			return
		default:
			if err := handlers.RunHTTPServer(ipAddress, port); err != nil {
				log.Errorln(err)
				exit(err)
			}
		}
	}
}

const errorReporter = "https://2bedd272821643e1b92c774d3fdf28e7@sentry.statping.com/2"
