package main

import (
	"fmt"
	"os"
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
		fmt.Printf("Statup v%v Exporting Static 'index.html' page...\n", VERSION)
		RenderBoxes()
		configs, _ = LoadConfig()
		setupMode = true
		mainProcess()
		time.Sleep(10 * time.Second)
		indexSource := ExportIndexHTML()
		SaveFile("./index.html", []byte(indexSource))
		fmt.Println("Exported Statup index page: 'index.html'")
	case "help":
		HelpEcho()
	case "update":
		fmt.Println("Sorry updating isn't available yet!")
	default:
		fmt.Println("Statup does not have the command you entered.")
		os.Exit(1)
	}
}

func HelpEcho() {
	fmt.Printf("Statup v%v - Statup.io\n", VERSION)
	fmt.Printf("A simple Application Status Monitor that is opensource and lightweight.\n")
	fmt.Printf("Commands:\n")
	fmt.Println("     statup                    - Main command to run Statup server")
	fmt.Println("     statup version            - Returns the current version of Statup")
	fmt.Println("     statup export             - Exports the index page as a static HTML for pushing")
	fmt.Println("                                 to Github Pages or your own FTP server. Export will")
	fmt.Println("                                 create 'index.html' in the current directory.")
	fmt.Println("     statup update             - Attempts to update to the latest version")
	fmt.Println("     statup help               - Shows the user basic information about Statup")
	fmt.Println("Give Statup a Star at https://github.com/hunterlong/statup")
}
