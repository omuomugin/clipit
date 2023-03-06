package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize configuration.",
	Long:  "Initialize configuration. This will create configuration file to store. Use default configuration, if not initialized.",
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set("editor", viper.GetString("editor"))
		if err := viper.SafeWriteConfig(); err != nil {
			cobra.CheckErr(err)
		}
	},
}

func init() {
	configCmd.AddCommand(configInitCmd)
}
