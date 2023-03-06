package cmd

import (
	"fmt"
	"github.com/omuomugin/clipit/support"
	"github.com/spf13/cobra"
	"os"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all saved commands.",
	Long:  "List all saved commands.",
	Run: func(cmd *cobra.Command, args []string) {
		filename := support.FilePathForCommands()
		content, err := os.ReadFile(filename)
		cobra.CheckErr(err)

		fmt.Print(string(content))
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
