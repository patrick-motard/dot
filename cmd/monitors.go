// Copyright © 2018 Patrick Motard <motard19@gmail.com>

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// screenCmd represents the screen command
var monitorsCmd = &cobra.Command{
	Use:   "monitors",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("monitors called")
	},
}

func init() {
	rootCmd.AddCommand(monitorsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// screenCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// screenCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
