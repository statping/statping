package main

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/statping/statping/utils"
	"io"
	"os"
	"os/exec"
)

var rootCmd = &cobra.Command{
	Use:     "statping",
	Version: VERSION,
	Short:   "A simple Application Status Monitor that is opensource and lightweight.",
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}

var updateCmd = &cobra.Command{
	Use:     "update",
	Example: "statping update",
	Short:   "Update to the latest version",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Infoln("Updating Statping to the latest version...")
		log.Infoln("curl -o- -L https://statping.com/install.sh | bash")
		curl, err := exec.LookPath("curl")
		if err != nil {
			return err
		}
		bash, err := exec.LookPath("bash")
		if err != nil {
			return err
		}

		ree := bytes.NewBuffer(nil)

		c1 := exec.Command(curl, "-o-", "-L", "https://statping.com/install.sh")
		c2 := exec.Command(bash)

		r, w := io.Pipe()
		c1.Stdout = w
		c2.Stdin = r

		var b2 bytes.Buffer
		c2.Stdout = &b2

		c1.Start()
		c2.Start()
		c1.Wait()
		w.Close()
		c2.Wait()
		io.Copy(ree, &b2)

		log.Infoln(ree.String())
		os.Exit(0)
		return nil
	},
}

var versionCmd = &cobra.Command{
	Use:     "version",
	Example: "statping version",
	Short:   "Print the version number of Statping",
	Run: func(cmd *cobra.Command, args []string) {
		if COMMIT != "" {
			fmt.Printf("%s (%s)\n", VERSION, COMMIT)
		} else {
			fmt.Printf("%s\n", VERSION)
		}
		os.Exit(0)
	},
}

var systemctlCmd = &cobra.Command{
	Use:     "systemctl [install/uninstall]",
	Example: "statping systemctl install",
	Short:   "Install or Uninstall systemctl services",
	RunE: func(cmd *cobra.Command, args []string) error {
		if args[1] == "install" {
			if len(args) < 3 {
				return errors.New("requires 'install <working_path> <port>'")
			}
		}
		port := utils.ToInt(args[2])
		if port == 0 {
			port = 80
		}
		if err := systemctlCli(args[1], args[0] == "uninstall", port); err != nil {
			return err
		}
		os.Exit(0)
		return nil
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("requires 'install <working_path>' or 'uninstall' as arguments")
		}
		return nil
	},
}

var assetsCmd = &cobra.Command{
	Use:     "assets",
	Example: "statping assets",
	Short:   "Dump all assets used locally to be edited",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := assetsCli(); err != nil {
			return err
		}
		os.Exit(0)
		return nil
	},
}

var exportCmd = &cobra.Command{
	Use:     "export",
	Example: "statping export",
	Short:   "Exports your Statping settings to a 'statping-export.json' file.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := exportCli(args); err != nil {
			return err
		}
		os.Exit(0)
		return nil
	},
}

var sassCmd = &cobra.Command{
	Use:     "sass",
	Example: "statping sass",
	Short:   "Compile .scss files into the css directory",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := sassCli(); err != nil {
			return err
		}
		os.Exit(0)
		return nil
	},
}

var envCmd = &cobra.Command{
	Use:     "env",
	Example: "statping env",
	Short:   "Return the configs that will be ran",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := envCli(); err != nil {
			return err
		}
		os.Exit(0)
		return nil
	},
}

var resetCmd = &cobra.Command{
	Use:     "reset",
	Example: "statping reset",
	Short:   "Start a fresh copy of Statping",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := resetCli(); err != nil {
			return err
		}
		os.Exit(0)
		return nil
	},
}

var onceCmd = &cobra.Command{
	Use:     "once",
	Example: "statping once",
	Short:   "Check all services 1 time and then quit",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := onceCli(); err != nil {
			return err
		}
		os.Exit(0)
		return nil
	},
}

var importCmd = &cobra.Command{
	Use:     "import [.json file]",
	Example: "statping import backup.json",
	Short:   "Imports settings from a previously saved JSON file.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := importCli(args); err != nil {
			return err
		}
		os.Exit(0)
		return nil
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires input file (.json)")
		}
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		exit(err)
	}
}
