// Copyright © 2018 Patrick Motard <motard19@gmail.com>

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
)

// printCmd represents the print command
var printCmd = &cobra.Command{
	Use:   "print",
	Short: "Prints settings from current_settings.toml if it exists, otherwise settings.toml",
	Long: `The default location for the .config file is found in the working directory of this project.
You can override the config file location by passing in --config.
This command will output the contents of the current config file.
`,
	Run: func(cmd *cobra.Command, args []string) {
		// for _, property := range viper.AllSettings() {
		// 	fmt.Println(property)
		// }
		// settings.PrettyPrint()
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// printCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// printCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
