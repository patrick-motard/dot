// Copyright © 2018 Patrick Motard <motard19@gmail.com>

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"os/exec"
	"strings"
)

// selectCmd represents the select command
var selectCmd = &cobra.Command{
	Use:   "select",
	Short: "Select an randr script to apply to your system.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			viper.Set("monitors.current", args[0])
		}
		runCurrent()
		viper.WriteConfig()
	},
}

func runCurrent() {
	location := viper.GetString("monitors.location")
	current := viper.GetString("monitors.current")
	fullPath := strings.Join([]string{location, current}, "/")
	fmt.Println(fullPath)
	arandrCmd := exec.Command("/bin/sh", fullPath)
	e := arandrCmd.Run()
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(1)
	}
}
func init() {
	screenCmd.AddCommand(selectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// selectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// selectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
