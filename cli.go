package main

import (
	"fmt"
	"github.com/hunterlong/statup/core"
	"github.com/hunterlong/statup/plugin"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"github.com/joho/godotenv"
	"math/rand"
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
	defer utils.DeleteFile("./.plugin_test.db")
	RenderBoxes()

	info := plug.GetInfo()
	fmt.Printf("\n" + BRAKER + "\n")
	fmt.Printf("    Plugin Name:          %v\n", info.Name)
	fmt.Printf("    Plugin Description:   %v\n", info.Description)
	fmt.Printf("    Plugin Routes:        %v\n", len(plug.Routes()))
	for k, r := range plug.Routes() {
		fmt.Printf("      - Route %v      - (%v) /%v \n", k+1, r.Method, r.URL)
	}

	// Function to create a new Core with example services, hits, failures, users, and default communications
	FakeSeed(plug)

	fmt.Println("\n" + BRAKER)
	fmt.Println(POINT + "Sending 'OnLoad(sqlbuilder.Database)'")
	core.OnLoad(core.DbSession)
	fmt.Println("\n" + BRAKER)
	fmt.Println(POINT + "Sending 'OnSuccess(Service)'")
	core.OnSuccess(core.SelectService(1))
	fmt.Println("\n" + BRAKER)
	fmt.Println(POINT + "Sending 'OnFailure(Service, FailureData)'")
	fakeFailD := core.FailureData{
		Issue: "No issue, just testing this plugin. This would include HTTP failure information though",
	}
	core.OnFailure(core.SelectService(1), fakeFailD)
	fmt.Println("\n" + BRAKER)
	fmt.Println(POINT + "Sending 'OnSettingsSaved(Core)'")
	fmt.Println(BRAKER)
	core.OnSettingsSaved(core.CoreApp)
	fmt.Println("\n" + BRAKER)
	fmt.Println(POINT + "Sending 'OnNewService(Service)'")
	core.OnNewService(core.SelectService(2))
	fmt.Println("\n" + BRAKER)
	fmt.Println(POINT + "Sending 'OnNewUser(User)'")
	user, _ := core.SelectUser(1)
	core.OnNewUser(user)
	fmt.Println("\n" + BRAKER)
	fmt.Println(POINT + "Sending 'OnUpdateService(Service)'")
	srv := core.SelectService(2)
	srv.Type = "http"
	srv.Domain = "https://yahoo.com"
	core.OnUpdateService(srv)
	fmt.Println("\n" + BRAKER)
	fmt.Println(POINT + "Sending 'OnDeletedService(Service)'")
	core.OnDeletedService(core.SelectService(1))
	fmt.Println("\n" + BRAKER)
}

func FakeSeed(plug plugin.PluginActions) {
	var err error
	core.CoreApp = core.NewCore()

	core.CoreApp.AllPlugins = []plugin.PluginActions{plug}

	fmt.Printf("\n" + BRAKER)

	fmt.Println("\nCreating a SQLite database for testing, will be deleted automatically...")
	sqlFake := sqlite.ConnectionURL{
		Database: "./.plugin_test.db",
	}
	core.DbSession, err = sqlite.Open(sqlFake)
	if err != nil {
		utils.Log(3, err)
	}
	up, _ := core.SqlBox.String("sqlite_up.sql")
	requests := strings.Split(up, ";")
	for _, request := range requests {
		_, err := core.DbSession.Exec(request)
		if err != nil {
			utils.Log(2, err)
		}
	}

	fmt.Println("Finished creating Test SQLite database")
	fmt.Println("Inserting example services into test database...")

	core.CoreApp.Name = "Plugin Test"
	core.CoreApp.Description = "This is a fake Core for testing your plugin"
	core.CoreApp.Domain = "http://localhost:8080"
	core.CoreApp.ApiSecret = "0x0x0x0x0"
	core.CoreApp.ApiKey = "abcdefg12345"

	fakeSrv := &core.Service{
		Name:   "Test Plugin Service",
		Domain: "https://google.com",
		Method: "GET",
	}
	fakeSrv.Create()

	fakeSrv2 := &core.Service{
		Name:   "Awesome Plugin Service",
		Domain: "https://netflix.com",
		Method: "GET",
	}
	fakeSrv2.Create()

	fakeUser := &core.User{
		Id:        6334,
		Username:  "Bulbasaur",
		Password:  "$2a$14$NzT/fLdE3f9iB1Eux2C84O6ZoPhI4NfY0Ke32qllCFo8pMTkUPZzy",
		Email:     "info@testdomain.com",
		Admin:     true,
		CreatedAt: time.Now(),
	}
	fakeUser.Create()

	fakeUser = &core.User{
		Id:        6335,
		Username:  "Billy",
		Password:  "$2a$14$NzT/fLdE3f9iB1Eux2C84O6ZoPhI4NfY0Ke32qllCFo8pMTkUPZzy",
		Email:     "info@awesome.com",
		CreatedAt: time.Now(),
	}
	fakeUser.Create()

	comm := &types.Communication{
		Id:     1,
		Method: "email",
	}
	core.Create(comm)

	comm2 := &types.Communication{
		Id:     2,
		Method: "slack",
	}
	core.Create(comm2)

	for i := 0; i <= 50; i++ {
		dd := core.HitData{
			Latency: rand.Float64(),
		}
		fakeSrv.CreateHit(dd)
		dd = core.HitData{
			Latency: rand.Float64(),
		}
		fakeSrv2.CreateHit(dd)
		fail := core.FailureData{
			Issue: "This is not an issue, but it would container HTTP response errors.",
		}
		fakeSrv.CreateFailure(fail)

		fail = core.FailureData{
			Issue: "HTTP Status Code 521 did not match 200",
		}
		fakeSrv2.CreateFailure(fail)
	}

	fmt.Println("Seeding example data is complete, running Plugin Tests")

}
