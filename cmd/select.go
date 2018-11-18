// Copyright © 2018 Patrick Motard <motard19@gmail.com>

package cmd

import (
	// "errors"
	"fmt"
	"github.com/manifoldco/promptui"
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
	Long: `DESCRIPTION:
Select an RandR script from a list or call it directly by name. Dot will
remember what you chose, even if you log out or reboot.

CALLING:

With no argument:

dot screen select
Use the arrow keys to navigate: ↓ ↑ → ←
? Pick one:
  ▸ 1080L-1440R
    1440
    1440-HDMI0-L-1440-DP4-R
    1440L-1080R


With argument:

dot screen select 1440L-1080R


NOTES:

Scripts are from 'monitors.location' in your config. I like to use ~/.screenlayout/

Each script is a shell script created by aRandR (a GUI for RandR).
RandR is a CLI for setting screen orientation, resolution, and rotation. With aRandR
you can visualy orient your monitors. Use aRandR to save the layout to a shell script
then call this script to apply it to your system.

Sometimes when you re-plugin monitors to your graphics card, it can
invalidate your layout made with aRandR. It's a real annoyance that I don't have a
fix for right now. You will have to recreate the layout and re-run 'dot screen select'
`,
	Args: cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		files, err := MonitorLocations()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		if len(args) == 0 {
			prompt := promptui.Select{
				Label: "Pick one",
				Items: files,
			}
			_, result, err := prompt.Run()
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			viper.Set("monitors.current", result)
		}
		if len(args) != 0 {
			viper.Set("monitors.current", args[0])
		}
		viper.WriteConfig()
		runCurrent()
	},
}

func runCurrent() {
	location := viper.GetString("monitors.location")
	current := viper.GetString("monitors.current")
	fullPath := strings.Join([]string{location, current}, "/")
	arandrCmd := exec.Command("/bin/sh", fullPath)
	out, err := arandrCmd.CombinedOutput()
	if err != nil {
		fmt.Println(fmt.Sprintf("Failed to run script %s", current))
		fmt.Println(out)
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
func init() {
	monitorsCmd.AddCommand(selectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// selectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// selectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
