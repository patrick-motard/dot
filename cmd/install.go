// Copyright © 2018 Patrick Motard <motard19@gmail.com>

package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "This command is a work in progress. Come back later. :)",
	Long: `Some ideas for this command:
- Allow users to install pacman packages and optionally 'save' the package to their list of packages they want to sync on other computers.
- Allow users to update Arch and AUR packages from the command line. Not sure what the point of this is yet.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("install called")
		// pac := exec.Command("/bin/sh", "-c", "sudo pacman -Syu --noconfirm")
		pac := exec.Command("/bin/sh", "-c", "echo hello world")
		out, err := pac.CombinedOutput()
		if err != nil {
			fmt.Println(string(out))
			fmt.Println(err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
