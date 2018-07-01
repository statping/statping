package main

import (
	"fmt"
	"github.com/hunterlong/statup/core"
	"github.com/hunterlong/statup/utils"
	"github.com/joho/godotenv"
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
	fmt.Println("     statup run                - Check all service 1 time and then quit")
	fmt.Println("     statup assets             - Export all assets used locally to be edited.")
	fmt.Println("     statup env                - Show all environment variables being used for Statup")
	fmt.Println("     statup export             - Exports the index page as a static HTML for pushing")
	fmt.Println("                                 to Github Pages or your own FTP server. Export will")
	fmt.Println("                                 create 'index.html' in the current directory.")
	fmt.Println("     statup update             - Attempts to update to the latest version")
	fmt.Println("     statup help               - Shows the user basic information about Statup")
	fmt.Println("Give Statup a Star at https://github.com/hunterlong/statup")
}
