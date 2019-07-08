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
	files, readErr := ioutil.ReadDir(Home + "/" + viper.GetString("displays.location"))
	var filenames []string
	for _, file := range files {
		filenames = append(filenames, file.Name())
	}
	return filenames, readErr
}

func init() {
	displaysCmd.AddCommand(displaysListCmd)
}
