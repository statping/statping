package main

import (
	"github.com/spf13/cobra"
	"github.com/statping/statping/utils"
)

var (
	ipAddress   string
	configFile  string
	verboseMode int
	port        int
)

func parseFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&ipAddress, "ip", "s", "0.0.0.0", "server run on host")
	utils.Params.BindPFlag("ip", cmd.PersistentFlags().Lookup("ip"))

	cmd.PersistentFlags().IntVarP(&port, "port", "p", 8080, "server port")
	utils.Params.BindPFlag("port", cmd.PersistentFlags().Lookup("port"))

	cmd.PersistentFlags().IntVarP(&verboseMode, "verbose", "v", 2, "verbose logging")
	utils.Params.BindPFlag("verbose", cmd.PersistentFlags().Lookup("verbose"))

	cmd.PersistentFlags().StringVarP(&configFile, "config", "c", utils.Directory+"/config.yml", "path to config.yml file")
	utils.Params.BindPFlag("config", cmd.PersistentFlags().Lookup("config"))
}
