package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func parseFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&ipAddress, "ip", "s", "0.0.0.0", "server port")
	viper.BindPFlag("ip", cmd.PersistentFlags().Lookup("ip"))

	cmd.PersistentFlags().IntVarP(&port, "port", "p", 8080, "server port")
	viper.BindPFlag("port", cmd.PersistentFlags().Lookup("port"))

	cmd.PersistentFlags().IntVarP(&verboseMode, "verbose", "v", 2, "server port")
	viper.BindPFlag("verbose", cmd.PersistentFlags().Lookup("verbose"))
}
