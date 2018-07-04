package main

import (
	"fmt"
	"github.com/hunterlong/statup/core"
	"github.com/hunterlong/statup/plugin"
	"github.com/hunterlong/statup/utils"
	"github.com/joho/godotenv"
	"strings"
	"time"
	"upper.io/db.v3/sqlite"
)

const (
	BRAKER = "=============================================================================="
	POINT  = "                     "
)

func CatchCLI(args []string) {
	switch args[1] {
	case "version":
		fmt.Printf("Statup v%v\n", VERSION)
	case "assets":
		RenderBoxes()
		core.CreateAllAssets()
	case "sass":
		core.CompileSASS()
	case "api":
		HelpEcho()
	case "test":
		cmd := args[2]
		switch cmd {
		case "plugins":
			LoadPlugins(true)
		}
	case "export":
		var err error
		fmt.Printf("Statup v%v Exporting Static 'index.html' page...\n", VERSION)
		RenderBoxes()
		core.Configs, err = core.LoadConfig()
		if err != nil {
			utils.Log(4, "config.yml file not found")
		}
		RunOnce()
		indexSource := core.ExportIndexHTML()
		err = core.SaveFile("./index.html", []byte(indexSource))
		if err != nil {
			utils.Log(4, err)
		}
		utils.Log(1, "Exported Statup index page: 'index.html'")
	case "help":
		HelpEcho()
	case "update":
		fmt.Println("Sorry updating isn't available yet!")
	case "run":
		utils.Log(1, "Running 1 time and saving to database...")
		RunOnce()
		fmt.Println("Check is complete.")
	case "env":
		fmt.Println("Statup Environment Variables")
		envs, err := godotenv.Read(".env")
		if err != nil {
			utils.Log(4, "No .env file found in current directory.")
		}
		for k, e := range envs {
			fmt.Printf("%v=%v\n", k, e)
		}
	default:
		utils.Log(3, "Statup does not have the command you entered.")
	}
}

func RunOnce() {
	var err error
	core.Configs, err = core.LoadConfig()
	if err != nil {
		utils.Log(4, "config.yml file not found")
	}
	err = core.DbConnection(core.Configs.Connection)
	if err != nil {
		utils.Log(4, err)
	}
	core.CoreApp, err = core.SelectCore()
	if err != nil {
		fmt.Println("Core database was not found, Statup is not setup yet.")
	}
	core.CoreApp.Services, err = core.SelectAllServices()
	if err != nil {
		utils.Log(4, err)
	}
	for _, s := range core.CoreApp.Services {
		out := s.Check()
		fmt.Printf("    Service %v | URL: %v | Latency: %0.0fms | Online: %v\n", out.Name, out.Domain, (out.Latency * 1000), out.Online)
	}
}

func HelpEcho() {
	fmt.Printf("Statup v%v - Statup.io\n", VERSION)
	fmt.Printf("A simple Application Status Monitor that is opensource and lightweight.\n")
	fmt.Printf("Commands:\n")
	fmt.Println("     statup                    - Main command to run Statup server")
	fmt.Println("     statup version            - Returns the current version of Statup")
	fmt.Println("     statup run                - Check all services 1 time and then quit")
	fmt.Println("     statup test plugins       - Test all plugins for required information")
	fmt.Println("     statup assets             - Dump all assets used locally to be edited.")
	fmt.Println("     statup env                - Show all environment variables being used for Statup")
	fmt.Println("     statup export             - Exports the index page as a static HTML for pushing")
	fmt.Println("     statup update             - Attempts to update to the latest version")
	fmt.Println("     statup help               - Shows the user basic information about Statup")
	fmt.Println("Give Statup a Star at https://github.com/hunterlong/statup")
}

func TestPlugin(plug plugin.PluginActions) {
	RenderBoxes()
	defer utils.DeleteFile("./.plugin_test.db")
	core.CoreApp.AllPlugins = []plugin.PluginActions{plug}
	info := plug.GetInfo()
	fmt.Printf("\n" + BRAKER + "\n")
	fmt.Printf("    Plugin Name:          %v\n", info.Name)
	fmt.Printf("    Plugin Description:   %v\n", info.Description)
	fmt.Printf("    Plugin Routes:        %v\n", len(plug.Routes()))
	for k, r := range plug.Routes() {
		fmt.Printf("      - Route %v      - (%v) /%v \n", k+1, r.Method, r.URL)
	}

	fmt.Printf("\n" + BRAKER)

	fakeSrv := &core.Service{
		Id:     56,
		Name:   "Test Plugin Service",
		Domain: "https://google.com",
	}

	fakeFailD := core.FailureData{
		Issue: "No issue, just testing this plugin.",
	}

	fakeCore := &core.Core{
		Name:        "Plugin Test",
		Description: "This is a fake Core for testing your plugin",
		ApiSecret:   "0x0x0x0x0",
		ApiKey:      "abcdefg12345",
		Services:    []*core.Service{fakeSrv},
	}

	fakeUser := &core.User{
		Id:        6334,
		Username:  "Bulbasaur",
		Password:  "$2a$14$NzT/fLdE3f9iB1Eux2C84O6ZoPhI4NfY0Ke32qllCFo8pMTkUPZzy",
		Email:     "info@testdomain.com",
		Admin:     true,
		CreatedAt: time.Now(),
	}

	fmt.Println("\nCreating a SQLite database for testing, will be deleted automatically...")
	sqlFake := sqlite.ConnectionURL{
		Database: "./.plugin_test.db",
	}
	fakeDb, err := sqlite.Open(sqlFake)
	if err != nil {
		utils.Log(3, err)
	}
	up, _ := core.SqlBox.String("sqlite_up.sql")
	requests := strings.Split(up, ";")
	for _, request := range requests {
		_, err := fakeDb.Exec(request)
		if err != nil {
			utils.Log(2, err)
		}
	}
	fmt.Println("Finished creating Test SQLite database, sending events.")

	fmt.Println("\n" + BRAKER)
	fmt.Println(POINT + "Sending 'OnLoad(sqlbuilder.Database)'")
	core.OnLoad(fakeDb)
	fmt.Println("\n" + BRAKER)
	fmt.Println(POINT + "Sending 'OnSuccess(Service)'")
	core.OnSuccess(fakeSrv)
	fmt.Println("\n" + BRAKER)
	fmt.Println(POINT + "Sending 'OnFailure(Service, FailureData)'")
	core.OnFailure(fakeSrv, fakeFailD)
	fmt.Println("\n" + BRAKER)
	fmt.Println(POINT + "Sending 'OnSettingsSaved(Core)'")
	fmt.Println(BRAKER)
	core.OnSettingsSaved(fakeCore)
	fmt.Println("\n" + BRAKER)
	fmt.Println(POINT + "Sending 'OnNewService(Service)'")
	core.OnNewService(fakeSrv)
	fmt.Println("\n" + BRAKER)
	fmt.Println(POINT + "Sending 'OnNewUser(User)'")
	core.OnNewUser(fakeUser)
	fmt.Println("\n" + BRAKER)
	fmt.Println(POINT + "Sending 'OnUpdateService(Service)'")
	core.OnUpdateService(fakeSrv)
	fmt.Println("\n" + BRAKER)
	fmt.Println(POINT + "Sending 'OnDeletedService(Service)'")
	core.OnDeletedService(fakeSrv)
	fmt.Println("\n" + BRAKER)

}

func FakeSeed() {

}
