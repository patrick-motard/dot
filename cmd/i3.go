// Copyright © 2018 Patrick Motard <motard19@gmail.com>

package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
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
	Short:     "This command is a work in progress. Come back later. :)",
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
	rootCmd.AddCommand(i3Cmd)
}
