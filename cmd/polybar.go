// Copyright © 2018 Patrick Motard <motard19@gmail.com>

package cmd

import (
	"fmt"
	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/randr"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/spf13/cobra"
	// "os"
)

var polybarCmd = &cobra.Command{
	Use:   "polybar",
	Short: "Loads polybar themes and bars.",
	Long:  "TODO: add long description",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You ran dot with the 'polybar' arguement.")
		main()
	},
}

func init() {
	rootCmd.AddCommand(polybarCmd)
}
func main() {
	X, _ := xgb.NewConn()
	err := randr.Init(X)
	if err != nil {
		log.Fatal(err)
	}
	root := xproto.Setup(X).DefaultScreen(X).Root
	resources, err := randr.GetScreenResources(X, root).Reply()
	if err != nil {
		log.Fatal(err)
	}
	for _, output := range resources.Outputs {
		info, err := randr.GetOutputInfo(X, output, 0).Reply()
		if err != nil {
			log.Fatal(err)
		}

		// fmt.Println(randr.ConnectionConnected)
		// fmt.Println(info.Connection)
		if info.Connection == randr.ConnectionConnected {
			fmt.Println(string(info.Name))
			crtc, _ := randr.GetCrtcInfo(X, info.Crtc, 0).Reply()
			fmt.Println(crtc.X)

			bestMode := info.Modes[0]
			for _, mode := range resources.Modes {
				if mode.Id == uint32(bestMode) {
					fmt.Printf("Width: %d, Height: %d, Name: %d\n", mode.Width, mode.Height, mode.Id)
				}
			}
		}
	}
	// fmt.Println("")

	// for _, outputs := range resources.Outputs {

	// 	fmt.Println(outputs)

	// }
	for _, crtc := range resources.Crtcs {
		info, err := randr.GetCrtcInfo(X, crtc, 0).Reply()
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(info.Help)
		// fmt.Println(string(info))
		fmt.Printf("X: %d, Y: %d, Width: %d, Height: %d\n",
			info.X, info.Y, info.Width, info.Height)
	}
}
