// Copyright © 2018 Patrick Motard <motard19@gmail.com>

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// soundCmd represents the sound command
var soundCmd = &cobra.Command{
	Use:   "sound",
	Short: "The sound command holds all commands needed for interacting with the systems audio.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sound called")
	},
}

func init() {
	rootCmd.AddCommand(soundCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// soundCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
