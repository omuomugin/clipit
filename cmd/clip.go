package cmd

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/omuomugin/clipit/support"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var clipCmd = &cobra.Command{
	Use:   "clip",
	Short: "Save snippet.",
	Long:  "Save snippet. If nothing is passed, prompt for inputting command will launch automatically.",
	Run: func(cmd *cobra.Command, args []string) {
		var command string

		if len(args) > 0 {
			command = strings.Join(args, " ")
		} else {
			prompt := promptui.Prompt{
				Label: "command to save",
			}
			result, _ := prompt.Run()
			command = result
		}

		filename := support.FilePathForCommands()
		w, errOpenFile := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		cobra.CheckErr(errOpenFile)
		defer w.Close()

		if _, err := fmt.Fprintln(w, command); err != nil {
			cobra.CheckErr(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(clipCmd)
}
