// Copyright © 2018 Patrick Motard <motard19@gmail.com>

package cmd

import (
	"github.com/spf13/cobra"
	// "strings"
)

// selectCmd represents the select command
var selectCmd = &cobra.Command{
	Use:   "select",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("select called")
		if len(args) == 0 {
			settings.Monitors.RunCurrent()
			return
		}
		settings.Monitors.SetCurrent(args[0])
		settings.Monitors.RunCurrent()
		settings.WriteSettings()
	},
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
