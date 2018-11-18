// Copyright © 2018 Patrick Motard <motard19@gmail.com>

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Output list of randr scripts on this system.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
		files, err := MonitorLocations()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		for _, file := range files {
			fmt.Println(file)
		}
	},
}

func MonitorLocations() ([]string, error) {
	files, readErr := ioutil.ReadDir(viper.GetString("monitors.location"))
	var filenames []string
	for _, file := range files {
		filenames = append(filenames, file.Name())
	}
	return filenames, readErr
}

func init() {
	monitorsCmd.AddCommand(listCmd)

	// fmt.Println(settings.Monitors.Location)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
