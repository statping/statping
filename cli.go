package main

import (
	"fmt"
	"github.com/hunterlong/statup/log"
	"github.com/joho/godotenv"
	"time"
)

func CatchCLI(args []string) {
	switch args[1] {
	case "version":
		fmt.Printf("Statup v%v\n", VERSION)
	case "assets":
		RenderBoxes()
		CreateAllAssets()
	case "sass":
		CompileSASS()
	case "export":
		var err error
		fmt.Printf("Statup v%v Exporting Static 'index.html' page...\n", VERSION)
		RenderBoxes()
		configs, err = LoadConfig()
		if err != nil {
			log.Send(3, "config.yml file not found")
		}
		setupMode = true
		mainProcess()
		time.Sleep(10 * time.Second)
		indexSource := ExportIndexHTML()
		err = SaveFile("./index.html", []byte(indexSource))
		if err != nil {
			log.Send(2, err)
		}
		log.Send(1, "Exported Statup index page: 'index.html'")
	case "help":
		HelpEcho()
	case "update":
		fmt.Println("Sorry updating isn't available yet!")
	case "run":
		log.Send(1, "Running 1 time and saving to database...")
		var err error
		configs, err = LoadConfig()
		if err != nil {
			log.Send(3, "config.yml file not found")
		}
		err = DbConnection(configs.Connection)
		if err != nil {
			log.Send(3, err)
		}
		core, err = SelectCore()
		if err != nil {
			fmt.Println("Core database was not found, Statup is not setup yet.")
		}
		services, err = SelectAllServices()
		if err != nil {
			log.Send(3, err)
		}
		for _, s := range services {
			out := s.Check()
			fmt.Printf("    Service %v | URL: %v | Latency: %0.0fms | Online: %v\n", out.Name, out.Domain, (out.Latency * 1000), out.Online)
		}
		fmt.Println("Check is complete.")
	case "env":
		fmt.Println("Statup Environment Variables")
		envs, err := godotenv.Read(".env")
		if err != nil {
			log.Send(3, "No .env file found in current directory.")
		}
		for k, e := range envs {
			fmt.Printf("%v=%v\n", k, e)
		}
	default:
		log.Send(3, "Statup does not have the command you entered.")
	}
}

func HelpEcho() {
	fmt.Printf("Statup v%v - Statup.io\n", VERSION)
	fmt.Printf("A simple Application Status Monitor that is opensource and lightweight.\n")
	fmt.Printf("Commands:\n")
	fmt.Println("     statup                    - Main command to run Statup server")
	fmt.Println("     statup version            - Returns the current version of Statup")
	fmt.Println("     statup run                - Check all service 1 time and then quit")
	fmt.Println("     statup env                - Show all environment variables being used for Statup")
	fmt.Println("     statup export             - Exports the index page as a static HTML for pushing")
	fmt.Println("                                 to Github Pages or your own FTP server. Export will")
	fmt.Println("                                 create 'index.html' in the current directory.")
	fmt.Println("     statup update             - Attempts to update to the latest version")
	fmt.Println("     statup help               - Shows the user basic information about Statup")
	fmt.Println("Give Statup a Star at https://github.com/hunterlong/statup")
}
