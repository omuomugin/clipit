package cmd

import (
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Commands for setting configuration.",
	Long:  "Commands for setting configuration.",
}

func init() {
	rootCmd.AddCommand(configCmd)
}
