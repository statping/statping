package main

import (
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
	verboseMode int
	port        int
	log         = utils.Log.WithField("type", "cmd")
	confgs      *configs.DbConfig
)

func init() {
	core.New(VERSION)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(assetsCmd)
	rootCmd.AddCommand(exportCmd)
	rootCmd.AddCommand(importCmd)
	rootCmd.AddCommand(sassCmd)
	rootCmd.AddCommand(onceCmd)
	rootCmd.AddCommand(envCmd)
	rootCmd.AddCommand(resetCmd)
	utils.InitCLI()
	parseFlags(rootCmd)
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
	Execute()
}

// main will run the Statping application
func start() {
	var err error
	go sigterm()

	if err := source.Assets(); err != nil {
		exit(err)
	}

	utils.VerboseMode = verboseMode

	if err := utils.InitLogs(); err != nil {
		log.Errorf("Statping Log Error: %v\n", err)
	}

	log.Info(fmt.Sprintf("Starting Statping v%s", VERSION))

	//if err := updateDisplay(); err != nil {
	//	log.Warnln(err)
	//}

	confgs, err = configs.LoadConfigs()
	if err != nil {
		log.Infoln("Starting in Setup Mode")
		if err := SetupMode(); err != nil {
			exit(err)
		}
	}

	if err = configs.ConnectConfigs(confgs, true); err != nil {
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

		if utils.Params.GetBool("SAMPLE_DATA") {
			if err := configs.TriggerSamples(); err != nil {
				exit(errors.Wrap(err, "error creating database"))
			}
		} else {
			if err := core.Samples(); err != nil {
				exit(errors.Wrap(err, "error added core details"))
			}
		}

	}

	if err = confgs.DatabaseChanges(); err != nil {
		exit(err)
	}

	if err := confgs.MigrateDatabase(); err != nil {
		exit(err)
	}

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
	if err := InitApp(); err != nil {
		return err
	}

	services.LoadServicesYaml()

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
