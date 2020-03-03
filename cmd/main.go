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
	"github.com/hunterlong/statping/utils"
	"github.com/pkg/errors"

	"flag"
	"fmt"
	"github.com/hunterlong/statping/core"
	"github.com/hunterlong/statping/handlers"
	"github.com/hunterlong/statping/source"
	"github.com/joho/godotenv"
	"os"
	"os/signal"
	"syscall"
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
)

func init() {
	core.VERSION = VERSION
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
	fmt.Printf("%+v", core.Configs())
	panic(err)
	//log.Fatalln(err)
	//os.Exit(2)
}

// main will run the Statping application
func main() {
	var err error
	go sigterm()
	parseFlags()
	loadDotEnvs()
	source.Assets()
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

	// check if DB_CONN was set, and load config from that
	autoConfigDb := utils.Getenv("DB_CONN", "").(string)
	if autoConfigDb != "" {
		log.Infof("Environment variable 'DB_CONN' was set to %s, loading configs from ENV.", autoConfigDb)
		if _, err := core.LoadUsingEnv(); err != nil {
			exit(err)
			return
		} else {
			afterConfigLoaded()
		}
	}

	// attempt to load config.yml file from current directory, if no file, then start in setup mode.
	_, err = core.LoadConfigFile(utils.Directory)
	if err != nil {
		log.Errorln(err)
		core.CoreApp.Setup = false
		writeAble, err := utils.DirWritable(utils.Directory)
		if err != nil {
			exit(err)
			return
		}
		if !writeAble {
			log.Fatalf("Statping does not have write permissions at: %v\nYou can change this directory by setting the STATPING_DIR environment variable to a dedicated path before starting.", utils.Directory)
			return
		}
		if err := handlers.RunHTTPServer(ipAddress, port); err != nil {
			log.Fatalln(err)
		}
	} else {
		afterConfigLoaded()
	}
}

func afterConfigLoaded() {
	if err := mainProcess(); err != nil {
		exit(err)
	}
}

// Close will gracefully stop the database connection, and log file
func Close() {
	utils.CloseLogs()
	core.CloseDB()
}

// sigterm will attempt to close the database connections gracefully
func sigterm() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	<-sigs
	Close()
	os.Exit(1)
}

// loadDotEnvs attempts to load database configs from a '.env' file in root directory
func loadDotEnvs() error {
	err := godotenv.Load(envFile)
	if err == nil {
		log.Infoln("Environment file '.env' Loaded")
	}
	return err
}

// mainProcess will initialize the Statping application and run the HTTP server
func mainProcess() error {
	dir := utils.Directory
	var err error
	err = core.CoreApp.Connect(core.Configs(), false, dir)
	if err != nil {
		log.Errorln(fmt.Sprintf("could not connect to database: %v", err))
		return err
	}

	if err := core.CoreApp.MigrateDatabase(); err != nil {
		return errors.Wrap(err, "database migration")
	}

	if err := core.CoreApp.ServicesFromEnvFile(); err != nil {
		errStr := "error 'SERVICE' environment variable"
		log.Errorln(errStr)
		return errors.Wrap(err, errStr)
	}

	if err := core.InitApp(); err != nil {
		return err
	}
	if core.CoreApp.Setup {
		if err := handlers.RunHTTPServer(ipAddress, port); err != nil {
			log.Fatalln(err)
			return errors.Wrap(err, "http server")
		}
	}
	return nil
}
