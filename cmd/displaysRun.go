// Copyright © 2018 Patrick Motard <motard19@gmail.com>

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Name string

var displaysRunCmd = &cobra.Command{
	Use:   "run",
	Short: "Apply an RandR script to your system by name.",
	Long:  `See 'dot display --help for more details.`,
	Run: func(cmd *cobra.Command, args []string) {
		if Name != "" {
			runErr := RunDisplaysScript(Name)
			if runErr != nil {
				fmt.Println(fmt.Sprintf("Failed to run script %s", Name))
				return
			}
			viper.Set("displays.current", Name)
			viper.WriteConfig()
			return
		}
		RunDisplaysScript(viper.GetString("displays.current"))
	},
}

func init() {
	displaysCmd.AddCommand(displaysRunCmd)
	displaysRunCmd.Flags().StringVarP(&Name, "name", "n", "", "Name of RandR script to run.")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// monitorsRunCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// monitorsRunCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
