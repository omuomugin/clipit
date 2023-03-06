package cmd

import (
	"github.com/omuomugin/clipit/support"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var configEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit configuration with selected editor.",
	Long:  "Edit configuration with selected editor.",
	Run: func(cmd *cobra.Command, args []string) {
		filename := support.FilePathForConfigs()
		command := viper.GetString("editor") + " " + filename
		support.Run(command, os.Stdin, os.Stdout)
	},
}

func init() {
	configCmd.AddCommand(configEditCmd)
}
