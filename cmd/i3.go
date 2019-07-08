// Copyright © 2018 Patrick Motard <motard19@gmail.com>

package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

var load string
var validArgs = []string{"load"}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var i3Cmd = &cobra.Command{
	Use:       "i3",
	Short:     "A brief description of your command",
	Long:      ``,
	ValidArgs: []string{"load"},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("Requires exactly 1 arg.")
		}
		return cobra.OnlyValidArgs(cmd, args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] != "load" {
			os.Exit(0)
		}
		files, err := ioutil.ReadDir(Home + "/.config/i3/")
		check(err)
		for _, f := range files {
			fmt.Println(f.Name())
		}
		fmt.Println(args)
	},
}

func init() {
	// i3Cmd.Flags().StringVarP(&load, "load", "", "", "load i3")
	rootCmd.AddCommand(i3Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// i3Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// i3Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
