package support

import (
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

func FilePathRoot() string {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	return filepath.Join(home, ".clipit")
}

func FilePathForCommands() string {
	home := FilePathRoot()
	return filepath.Join(home, ".commands")
}

func FilePathForConfigs() string {
	home := FilePathRoot()
	return filepath.Join(home, ".config.toml")
}
