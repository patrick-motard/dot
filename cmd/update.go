// Copyright © 2018 Patrick Motard <motard19@gmail.com>

package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "This command is a work in progress. Come back later. :)",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update called")

		cmdArgs := []string{"-c", "sudo pacman -Syu --noconfirm"}
		updatePacman := exec.Command("/bin/sh", cmdArgs...)
		out, err := updatePacman.CombinedOutput()
		fmt.Println(string(out))

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
