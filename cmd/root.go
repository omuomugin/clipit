package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "clipit",
	Short: "Clipit is simple CLI snippet manager.",
	Long:  "Clipit is simple CLI snippet manager.",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
