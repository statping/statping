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

	"flag"
	"fmt"
	"github.com/hunterlong/statping/core"
	"github.com/hunterlong/statping/handlers"
	"github.com/hunterlong/statping/plugin"
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
func parseFlags() {
	flag.StringVar(&ipAddress, "ip", "0.0.0.0", "IP address to run the Statping HTTP server")
	flag.StringVar(&envFile, "env", "", "IP address to run the Statping HTTP server")
	flag.IntVar(&port, "port", 8080, "Port to run the HTTP server")
	flag.IntVar(&verboseMode, "verbose", 2, "Run in verbose mode to see detailed logs (1 - 4)")
	flag.Parse()

	if os.Getenv("PORT") != "" {
		port = int(utils.ToInt(os.Getenv("PORT")))
	}
	if os.Getenv("IP") != "" {
		ipAddress = os.Getenv("IP")
	}
	if os.Getenv("VERBOSE") != "" {
		verboseMode = int(utils.ToInt(os.Getenv("VERBOSE")))
	}
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
			}
			fmt.Println(err)
			os.Exit(1)
		}
	}
	log.Info(fmt.Sprintf("Starting Statping v%v", VERSION))
	updateDisplay()

	configs, err := core.LoadConfigFile(utils.Directory)
	if err != nil {
		log.Errorln(err)
		core.SetupMode = true
		writeAble, err := utils.DirWritable(utils.Directory)
		if err != nil {
			log.Fatalln(err)
		}
		if !writeAble {
			log.Fatalf("Statping does not have write permissions at: %v\nYou can change this directory by setting the STATPING_DIR environment variable to a dedicated path before starting.", utils.Directory)
		}
		if err := handlers.RunHTTPServer(ipAddress, port); err != nil {
			log.Fatalln(err)
		}
	}
	core.CoreApp.Config = configs
	if err := mainProcess(); err != nil {
		log.Fatalln(err)
	}
}

// Close will gracefully stop the database connection, and log file
func Close() {
	core.CloseDB()
	utils.CloseLogs()
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
	err = core.CoreApp.Connect(false, dir)
	if err != nil {
		log.Errorln(fmt.Sprintf("could not connect to database: %v", err))
		return err
	}
	core.CoreApp.MigrateDatabase()
	core.InitApp()
	if !core.SetupMode {
		plugin.LoadPlugins()
		if err := handlers.RunHTTPServer(ipAddress, port); err != nil {
			log.Fatalln(err)
		}
	}
	return err
}
