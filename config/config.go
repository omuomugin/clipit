package config

import (
	"github.com/omuomugin/clipit/support"
	"github.com/spf13/viper"
)

type Config struct {
	Editor string `toml:"editor"`
}

func Load() {
	filename := support.FilePathRoot()
	viper.AddConfigPath(filename)
	viper.SetConfigName(".config")
	viper.SetConfigType("toml")

	viper.SetDefault("editor", "vi")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
		} else {
			// Config file was found but another error was produced
		}
	}
}
