// Copyright © 2018 Patrick Motard <motard19@gmail.com>

package cmd

import (
	"fmt"
	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/randr"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"os/exec"
	"sort"
	// "time"
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

type display struct {
	name string // example: DP-4 or HDMI-1
	// Position is where the display is relative to other displays on the screen.
	// Screens are comprised of one or more displays.
	xposition int16  // the x coordinate of the display on the screen
	yposition int16  // the y coordinate of the display on the screen
	xres      uint16 // The ideal x resolution.
	yres      uint16 // The idea y resolution.
	primary   bool   // Whether or not the display is the primary (main) display.
	active    bool
}

var theme string

func init() {
	polybarCmd.Flags().StringVarP(&theme, "theme", "t", "", "Specify theme by name. This will persist.")
	viper.BindPFlag("Polybar.Theme", polybarCmd.Flags().Lookup("theme"))
	rootCmd.AddCommand(polybarCmd)
}
func main() {
	// save the new theme if it is set
	// viper.WriteConfig()
	fmt.Println(Config.Polybar.Theme)
	fmt.Println(theme)
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
	primaryOutput, _ := randr.GetOutputPrimary(X, root).Reply()
	var displays []display
	for _, output := range resources.Outputs {
		info, err := randr.GetOutputInfo(X, output, 0).Reply()
		if err != nil {
			log.Fatal(err)
		}
		if info.Connection == randr.ConnectionConnected {
			d := display{
				name: string(info.Name),
			}
			crtc, err := randr.GetCrtcInfo(X, info.Crtc, 0).Reply()
			if err != nil {
				// log.Fatal("Failed to get CRTC info", err)
				// "BadCrtc" happens when attempting to get
				// a crtc for an output is disabled (inactive).
				// TODO: figure out a better way to identify active vs inactive

				d.active = false
			} else {
				d.active = true
				d.xposition = crtc.X
				d.yposition = crtc.Y
			}

			if output == primaryOutput.Output {
				d.primary = true
			} else {
				d.primary = false
			}
			bestMode := info.Modes[0]
			for _, mode := range resources.Modes {
				if mode.Id == uint32(bestMode) {
					d.xres = mode.Width
					d.yres = mode.Height
				}
			}
			displays = append(displays, d)
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
	var polybarEnvVars []string
	for i, d := range displays {
		// skip inactive monitors
		if !d.active {
			continue
		}
		if d.primary {
			s := fmt.Sprintf("MONITOR_MAIN=%s", d.name)
			polybarEnvVars = append(polybarEnvVars, s)
		} else if i == 0 {
			s := fmt.Sprintf("MONITOR_LEFT=%s", d.name)
			polybarEnvVars = append(polybarEnvVars, s)
		} else if i == 1 || i == 2 {
			s := fmt.Sprintf("MONITOR_RIGHT=%s", d.name)
			polybarEnvVars = append(polybarEnvVars, s)
		}

	}

	fmt.Println(Config.Polybar.Theme)
	// start polybar
	themePath := fmt.Sprintf("polybar_theme=/home/han/.config/polybar/%s/config", Config.Polybar.Theme)
	polybarEnvVars = append(polybarEnvVars, themePath)
	fmt.Println(polybarEnvVars)
	newEnv := append(os.Environ(), polybarEnvVars...)
	var theme Theme
	for _, t := range Config.Polybar.Themes {
		if Config.Polybar.Theme == t.Name {
			theme = t
			// fmt.Println("hello")
		}
	}
	if theme.Name == "" {
		log.Error("Theme not found, exiting.")
		os.Exit(1)
	}
	bars := theme.Bars

	// bars := []string{
	// 	"main.top.middle",
	// 	"left.top.middle",
	// 	"right.top.middle",
	// }
	// fmt.Println()
	// go runStuff(newEnv, bars)
	// time.Sleep(5 * time.Second)
	// fmt.Println("main terminated")
	// }

	// func runStuff(envs, bars []string) {

	// fmt.Printf("%+v\n", &Config)
	for _, bar := range bars {
		polybar(newEnv, bar)
	}
}

// // var done chan string
// done := make(chan string, 3)
// go runAllPolybar(newEnv, bars, done)
// for {
// 	v, ok := <-done
// 	if ok == false {
// 		break
// 	}
// 	fmt.Println("Received: ", v, ok)
// }
// // hack to allow goroutines to start
// // TODO: replace with channels
// // time.Sleep(1 * time.Second)
// }

func runAllPolybar(envs, bars []string, ch chan string) {
	for _, bar := range bars {
		fmt.Println("Got here")
		ch <- polybar(envs, bar)
	}
	close(ch)
}

func polybar(env []string, bar string) string {
	log.Printf("Starting bar %s", bar)
	s := fmt.Sprintf("polybar -r %s", bar)
	cmd := exec.Command("bash", "-c", s)
	cmd.Env = env
	cmd.Start()
	// out, err := cmd.CombinedOutput()
	// if err != nil {
	// 	// log.Fatal(err)
	// 	return err.Error()
	// 	// fmt.Println(err)
	// }
	// cmd.Wait()
	return fmt.Sprintf("Finished bar %s", bar)
	// return string(out)
	// fmt.Println(string(out))
	// done <- string(out)
}
