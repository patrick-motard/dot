// Copyright © 2018 Patrick Motard <motard19@gmail.com>

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
