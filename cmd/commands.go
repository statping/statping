package main

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "statping",
	Short: "A simple Application Status Monitor that is opensource and lightweight.",
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Statping",
	Run: func(cmd *cobra.Command, args []string) {
		if COMMIT != "" {
			fmt.Printf("%s (%s)\n", VERSION, COMMIT)
		} else {
			fmt.Printf("%s\n", VERSION)
		}
		os.Exit(0)
	},
}

var assetsCmd = &cobra.Command{
	Use:   "assets",
	Short: "Dump all assets used locally to be edited",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := assetsCli(); err != nil {
			return err
		}
		os.Exit(0)
		return nil
	},
}

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Exports your Statping settings to a 'statping-export.json' file.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := exportCli(args); err != nil {
			return err
		}
		os.Exit(0)
		return nil
	},
}

var sassCmd = &cobra.Command{
	Use:   "sass",
	Short: "Compile .scss files into the css directory",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := sassCli(); err != nil {
			return err
		}
		os.Exit(0)
		return nil
	},
}

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Return the configs that will be ran",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := envCli(); err != nil {
			return err
		}
		os.Exit(0)
		return nil
	},
}

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Start a fresh copy of Statping",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := resetCli(); err != nil {
			return err
		}
		os.Exit(0)
		return nil
	},
}

var onceCmd = &cobra.Command{
	Use:   "once",
	Short: "Check all services 1 time and then quit",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := onceCli(); err != nil {
			return err
		}
		os.Exit(0)
		return nil
	},
}

var importCmd = &cobra.Command{
	Use:   "import [.json file]",
	Short: "Imports settings from a previously saved JSON file.",
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
