package cmd

import (
	"github.com/omuomugin/clipit/support"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit commands with selected editor.",
	Long:  "Edit commands with selected editor.",
	Run: func(cmd *cobra.Command, args []string) {
		filename := support.FilePathForCommands()
		command := viper.GetString("editor") + " " + filename
		support.Run(command, os.Stdin, os.Stdout)
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
