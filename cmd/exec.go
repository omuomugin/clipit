package cmd

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/omuomugin/clipit/snippet"
	"github.com/omuomugin/clipit/support"
	"github.com/spf13/cobra"
	"os"
	"regexp"
	"strings"
)

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Execute passed command.",
	Long:  "Execute passed command. Commands can include variables. e.g <key>, <key=A>, <key=A|B|C>.",
	Run: func(cmd *cobra.Command, args []string) {
		var command string
		if len(args) > 0 {
			command = strings.Join(args, " ")
		} else {
			prompt := promptui.Prompt{
				Label: "command to exec",
			}
			result, _ := prompt.Run()
			command = result
		}

		// command variables will have 2 types
		// <key=A>, <key> -> launch prompt with default value if set.
		// <key=A|B|C> -> launch select prompt with choices of A, B and C.
		r, _ := regexp.Compile(`<(\S.+?\S)>`)

		// this will break up into 2 maps
		// e.g.
		// [[<key=A> (original), key=A], [<key> (original), key]]
		params := r.FindAllStringSubmatch(command, -1)

		var variables []snippet.Variable
		for _, p := range params {
			// this will break up into multiple variables
			// e.g.
			// key -> [], key=A -> [key, A], key=A|B|C -> [key, A|B|C]
			original, extractedVariable := p[0], p[1]

			splitted := strings.Split(extractedVariable, "=")
			if len(splitted) == 1 {
				variables = append(variables, snippet.Variable{
					Key:     original,
					Choices: []string{},
				})
			} else if len(splitted) == 2 {
				variables = append(variables, snippet.Variable{
					Key: original,
					// A|B|C -> [A, B, C]
					Choices: strings.Split(splitted[1], "|"),
				})
			} else {
				os.Exit(1)
			}
		}

		for _, variable := range variables {
			if len(variable.Choices) == 0 {
				prompt := promptui.Prompt{
					Label: fmt.Sprintf("Variable for %v", variable.Key),
				}
				result, _ := prompt.Run()
				command = strings.Replace(command, variable.Key, result, 1)
			} else if len(variable.Choices) == 1 {
				prompt := promptui.Prompt{
					Label:   fmt.Sprintf("Variable for %v", variable.Key),
					Default: variable.Choices[0],
				}
				result, _ := prompt.Run()
				command = strings.Replace(command, variable.Key, result, 1)
			} else {
				prompt := promptui.Select{
					Label: fmt.Sprintf("Variable for %v", variable.Key),
					Items: variable.Choices,
				}
				_, result, _ := prompt.Run()
				command = strings.Replace(command, variable.Key, result, 1)
			}
		}

		fmt.Printf("[INFO] Executed command : %v\n", command)

		support.Run(command, os.Stdin, os.Stdout)
	},
}

func init() {
	rootCmd.AddCommand(execCmd)
}
