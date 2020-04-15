package main

import (
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"github.com/statping/statping/database"
	"github.com/statping/statping/handlers"
	"github.com/statping/statping/notifiers"
	"github.com/statping/statping/source"
	"github.com/statping/statping/types/configs"
	"github.com/statping/statping/types/core"
	"github.com/statping/statping/types/services"
	"github.com/statping/statping/utils"
	"os"
	"os/signal"
	"syscall"
)

var (
	// VERSION stores the current version of Statping
	VERSION string
	// COMMIT stores the git commit hash for this version of Statping
	COMMIT    string
	ipAddress string
	//grpcPort    int
	envFile     string
	verboseMode int
	port        int
	log         = utils.Log.WithField("type", "cmd")
	confgs      *configs.DbConfig
)

// parseFlags will parse the application flags
// -ip = 0.0.0.0 IP address for outgoing HTTP server
// -port = 8080 Port number for outgoing HTTP server
// environment variables WILL overwrite flags
func parseFlags() {
	envPort := utils.Getenv("PORT", 8080).(int)
	envIpAddress := utils.Getenv("IP", "0.0.0.0").(string)
	envVerbose := utils.Getenv("VERBOSE", 2).(int)
	//envGrpcPort := utils.Getenv("GRPC_PORT", 0).(int)

	flag.StringVar(&ipAddress, "ip", envIpAddress, "IP address to run the Statping HTTP server")
	flag.StringVar(&envFile, "env", "", "IP address to run the Statping HTTP server")
	flag.IntVar(&port, "port", envPort, "Port to run the HTTP server")
	//flag.IntVar(&grpcPort, "grpc", envGrpcPort, "Port to run the gRPC server")
	flag.IntVar(&verboseMode, "verbose", envVerbose, "Run in verbose mode to see detailed logs (1 - 4)")
	flag.Parse()
}

func init() {
	core.New(VERSION)
}

// exit will return an error and return an exit code 1 due to this error
func exit(err error) {
	utils.SentryErr(err)
	Close()
	log.Fatalln(err)
}

// Close will gracefully stop the database connection, and log file
func Close() {
	utils.CloseLogs()
	confgs.Close()
	fmt.Println("Shutting down Statping")
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
	log.Info(fmt.Sprintf("Starting Statping v%s", VERSION))

	if err := updateDisplay(); err != nil {
		log.Warnln(err)
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

		if utils.Getenv("SAMPLE_DATA", true).(bool) {
			if err := configs.TriggerSamples(); err != nil {
				exit(errors.Wrap(err, "error creating database"))
			}
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

func SetupMode() error {
	return handlers.RunHTTPServer(ipAddress, port)
}

// sigterm will attempt to close the database connections gracefully
func sigterm() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
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

// InitApp will start the Statping instance with a valid database connection
// This function will gather all services in database, add/init Notifiers,
// and start the database cleanup routine
func InitApp() error {
	if _, err := core.Select(); err != nil {
		return err
	}
	if _, err := services.SelectAllServices(true); err != nil {
		return err
	}
	go services.CheckServices()
	notifiers.InitNotifiers()
	go database.Maintenance()
	utils.SentryInit(&VERSION, core.App.AllowReports.Bool)
	core.App.Setup = true
	core.App.Started = utils.Now()
	return nil
}
