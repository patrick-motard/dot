// Copyright © 2018 Patrick Motard <motard19@gmail.com>

package cmd

import (
	// "errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/patrick-motard/rofigo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var _rofi bool

// selectCmd represents the select command
var displaysSelectCmd = &cobra.Command{
	Use:   "select",
	Short: "Interactively select an RandR script from list apply it to your system.",
	Long: `Interactively select an RandR script from list apply it to your system.
You can select the display via rofi by setting the --rofi/-r flag.

Example (in command line):
dot screen select
Use the arrow keys to navigate: ↓ ↑ → ←
? Pick one:
  ▸ 1080L-1440R
    1440
    1440-HDMI0-L-1440-DP4-R
    1440L-1080R


Dot will remember what you chose, even if you log out or reboot.

`,
	Args: cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		files, err := DisplaysLocation()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		var selection string
		if _rofi == true {
			v := rofigo.New("xrandr script", files...)
			v.Show()
			selection = v.Selection
		} else {
			prompt := promptui.Select{
				Label: "Pick one",
				Items: files,
			}
			_, selection, err = prompt.Run()
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
		}
		runErr := RunDisplaysScript(selection)
		if runErr != nil {
			current := viper.GetString("displays.current")
			fmt.Println(fmt.Sprintf("Failed to run script %s", selection))
			fmt.Println(runErr.Error())
			fmt.Println(fmt.Sprintf("Reverting to previous script: %s", current))
			RunDisplaysScript(current)
			os.Exit(1)
		}
		viper.Set("displays.current", selection)
		viper.WriteConfig()
	},
}

func RunDisplaysScript(scriptName string) error {
	location := viper.GetString("displays.location")
	fullPath := strings.Join([]string{Home, "/", location, scriptName}, "/")
	arandrCmd := exec.Command("/bin/sh", fullPath)
	out, err := arandrCmd.CombinedOutput()
	fmt.Println(string(out))
	return err
}
func init() {
	displaysCmd.AddCommand(displaysSelectCmd)
	displaysSelectCmd.Flags().BoolVarP(&_rofi, "rofi", "r", false, "Use rofi to select display configuration.")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// selectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// selectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
