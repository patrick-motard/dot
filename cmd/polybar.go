// Copyright © 2018 Patrick Motard <motard19@gmail.com>

package cmd

import (
	"fmt"
	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/randr"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"sort"
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

type Display struct {
	name      string
	xposition int16
	yposition int16
	xres      int16
	yres      int16
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
	// fmt.Printf("OUTPUT!! %+v\n", root)

	y, _ := randr.GetOutputPrimary(X, root).Reply()
	fmt.Printf("OUTPUT!! %+v\n", y)

	// fmt.Printf("%+v\n", resources)
	var displays []Display
	for _, output := range resources.Outputs {
		info, err := randr.GetOutputInfo(X, output, 0).Reply()
		if err != nil {
			log.Fatal(err)
		}

		// fmt.Printf("OUTPUT!! %+v\n", info)
		// fmt.Println(randr.ConnectionConnected)
		// fmt.Println(info.Connection)
		if info.Connection == randr.ConnectionConnected {
			fmt.Printf("OUTPUT INFO: \n%+v\n\n", info)
			// fmt.Println(string(info.Name))
			crtc, _ := randr.GetCrtcInfo(X, info.Crtc, 0).Reply()
			fmt.Printf("%+v\n", crtc)
			display := Display{
				name:      string(info.Name),
				xposition: crtc.X,
				yposition: crtc.Y,
			}
			displays = append(displays, display)

			// fmt.Printf("%+v\n", display)
			// fmt.Println(crtc.X)

			// bestMode := info.Modes[0]
			// for _, mode := range resources.Modes {
			// 	if mode.Id == uint32(bestMode) {
			// 		fmt.Printf("%+v\n", mode)
			// 		fmt.Printf("Width: %d, Height: %d, Name: %d\n", mode.Width, mode.Height, mode.Id)
			// 	}
			// }
		}
	}
	// for _, crtc := range resources.Crtcs {
	// 	info, err := randr.GetCrtcInfo(X, crtc, 0).Reply()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	// fmt.Println(info.Help)
	// 	// fmt.Println(string(info))
	// 	fmt.Printf("X: %d, Y: %d, Width: %d, Height: %d, Status: %d\n",
	// 		info.X, info.Y, info.Width, info.Height, info.Status)
	// }
	sort.Slice(displays, func(i, j int) bool {
		return displays[i].xposition < displays[j].xposition
	})
	fmt.Printf("%+v\n", displays)

	// kill polybar
	cmd := exec.Command("sh", "-c", "killall -q polybar")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Infoln("Failed to kill polybar")
	}
	fmt.Println(string(out))

	// start polybar
	polybarEnvVars := []string{"MONITOR_MAIN=DP-4", "polybar_theme=/home/han/.config/polybar/nord/config"}
	newEnv := append(os.Environ(), polybarEnvVars...)
	cmd = exec.Command("bash", "-c", "polybar -r main.top.middle")
	cmd.Env = newEnv
	out, err = cmd.CombinedOutput()
	if err != nil {
		// log.Fatal(err)
		fmt.Println(err)
	}
	fmt.Println(string(out))
}
