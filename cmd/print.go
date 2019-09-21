// Copyright © 2018 Patrick Motard <motard19@gmail.com>

package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// printCmd represents the print command
var printCmd = &cobra.Command{
	Use:   "print",
	Short: "Prints settings from current_settings.yml if it exists, otherwise settings.toml",
	Long: `The default location for the .config file is found in the working directory of this project.
You can override the config file location by passing in --config.
This command will output the contents of the current config file.
`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info(viper.ConfigFileUsed())
		dat, err := ioutil.ReadFile(viper.ConfigFileUsed())
		if err != nil {
			log.Error(err)
			os.Exit(1)
		}
		fmt.Print(string(dat))
	},
}

func init() {
	rootCmd.AddCommand(printCmd)
}
