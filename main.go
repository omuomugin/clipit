package main

import (
	"github.com/omuomugin/clipit/cmd"
	"github.com/omuomugin/clipit/config"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

func main() {
	// start up
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	filename := filepath.Join(home, ".clipit")
	_ = os.MkdirAll(filename, os.ModePerm)

	config.Load()
	cmd.Execute()
}
