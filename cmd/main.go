// Statping
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/hunterlong/statping
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
	"os"
	"os/signal"
	"syscall"

	"github.com/hunterlong/statping/source"

	"github.com/hunterlong/statping/database"
	"github.com/hunterlong/statping/handlers"
	"github.com/hunterlong/statping/types/configs"
	"github.com/hunterlong/statping/types/core"
	"github.com/hunterlong/statping/types/services"
	"github.com/hunterlong/statping/utils"
	"github.com/pkg/errors"
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
)

func init() {

}

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
	panic(err)
	//log.Fatalln(err)
	//os.Exit(2)
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

	c, err := configs.LoadConfigs()
	if err != nil {
		if err := SetupMode(); err != nil {
			exit(err)
		}
	}

	if err = configs.ConnectConfigs(c, true); err != nil {
		exit(err)
	}

	if err := c.MigrateDatabase(); err != nil {
		exit(err)
	}

	if err := mainProcess(); err != nil {
		exit(err)
	}
}

// Close will gracefully stop the database connection, and log file
func Close() {
	utils.CloseLogs()
	database.Close()
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
	os.Exit(1)
}

// mainProcess will initialize the Statping application and run the HTTP server
func mainProcess() error {
	if err := services.ServicesFromEnvFile(); err != nil {
		errStr := "error 'SERVICE' environment variable"
		log.Errorln(errStr)
		return errors.Wrap(err, errStr)
	}

	if err := core.InitApp(); err != nil {
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
