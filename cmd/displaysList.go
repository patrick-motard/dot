// Copyright © 2018 Patrick Motard <motard19@gmail.com>

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
)

var displaysListCmd = &cobra.Command{
	Use:   "list",
	Short: "Output list of RandR scripts on this system.",
	Long:  `See 'dot display --help for more details.`,
	Run: func(cmd *cobra.Command, args []string) {
		files, err := DisplaysLocation()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		for _, file := range files {
			fmt.Println(file)
		}
	},
}

func DisplaysLocation() ([]string, error) {
	files, readErr := ioutil.ReadDir(viper.GetString("displays.location"))
	var filenames []string
	for _, file := range files {
		filenames = append(filenames, file.Name())
	}
	return filenames, readErr
}

func init() {
	displaysCmd.AddCommand(displaysListCmd)

	// fmt.Println(settings.Monitors.Location)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
