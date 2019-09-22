// Copyright © 2018 Patrick Motard <motard19@gmail.com>

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var _leftDisplay, _rightDisplay, _primaryDisplay bool
var displaysCmd = &cobra.Command{
	Use:   "displays",
	Short: "All commands for interacting with your displays",
	Long: `Description: The 'display' command contains all the commands for interacting with your display displays.

Scripts are from 'displays.location' in your config. I like to use ~/.screenlayout/

Each script is a shell script created by aRandR (a GUI for RandR).
RandR is a CLI for setting screen orientation, resolution, and rotation. With aRandR
you can visualy orient your displays. Use aRandR to save the layout to a shell script
then call this script to apply it to your system.

Sometimes when you re-plugin displays to your graphics card, it can
invalidate your layout made with aRandR. It's a real annoyance that I don't have a
fix for right now. You will have to recreate the layout and re-run 'dot screen select --rofi'`,
	Run: func(cmd *cobra.Command, args []string) {
		ds := displays{}
		if _leftDisplay {
			fmt.Println(ds.getLeft().name)
		}
		if _rightDisplay {
			fmt.Println(ds.getRight().name)
		}
		if _primaryDisplay {
			fmt.Println(ds.getPrimary().name)
		}
	},
}

func init() {
	rootCmd.AddCommand(displaysCmd)
	displaysCmd.Flags().BoolVarP(&_leftDisplay, "get-left", "l", false, "Get the name of your left display.")
	displaysCmd.Flags().BoolVarP(&_rightDisplay, "get-right", "r", false, "Get the name of your right display.")
	displaysCmd.Flags().BoolVarP(&_primaryDisplay, "get-primary", "p", false, "Get the name of your primary display.")
}
