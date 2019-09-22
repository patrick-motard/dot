// Copyright © 2018 Patrick Motard <motard19@gmail.com>

package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// --init/-i: Create a new config in $XDG_CONFIG_HOME or ~/.config/, error if it already exists.
// --print/-p: Print the current config file, and it's location.

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Interact with your dot config file.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if _print {
			print()
		} else if _copy {
			copyNewConfig()
		} else {
			cmd.Help()
		}
	},
}

var _print bool
var _copy bool

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().BoolVarP(&_print, "print", "p", false, "Print the current settings file to console.")
	configCmd.Flags().BoolVarP(&_copy, "create", "c", false, "Create a new config file.")
}

// Prints current_settings.yml to console.
func print() {
	log.Info(viper.ConfigFileUsed())
	dat, err := ioutil.ReadFile(viper.ConfigFileUsed())
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	fmt.Print(string(dat))
}

// Creates a new config (current_settings.yml) in $XDG_CONFIG_HOME or ~/.config/, error if it already exists.
func copyNewConfig() {
	path := os.Getenv("XDG_CONFIG_HOME")
	fmt.Println(path)
	// create path for config file
	// error/exit if file already exists
	// copy config file from project to path
	// print where it was written to
}
