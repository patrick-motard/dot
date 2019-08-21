// Copyright © 2018 Patrick Motard <motard19@gmail.com>
// TODO: add i3wm dimensions to bar settings

package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/randr"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/patrick-motard/rofigo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.i3wm.org/i3"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strings"
	"sync"
	// "time"
)

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

var (
	_theme                 string
	_list                  bool
	_select                bool
	themeIsValid           bool
	InstalledPolybarThemes []string
	FullThemePath          string
	FullThemesPath         string
)

var polybarCmd = &cobra.Command{
	Use:   "polybar",
	Short: "Loads polybar themes and bars.",
	Long:  "TODO: add long description",
	Run: func(cmd *cobra.Command, args []string) {
		//TODO: fix this function. It's confusing, terribly organized, etc.

		FullThemesPath = Home + "/" + Config.Polybar.ThemesDirectory
		findThemes()
		if _list == true {
			listThemes()
			os.Exit(0)
		}
		if _select == true {
			v := rofigo.New("Select Polybar theme:", InstalledPolybarThemes...)
			v.Show()
			log.Infof("You selected: %s", v.Selection)

			_theme = v.Selection
		}

		if _theme == "" {
			_theme = Config.Polybar.Theme
			log.Infof("No theme specified, reloading default: \"%s\"", _theme)
		}
		themeIsValid = validateTheme()
		if themeIsValid == false {
			log.Fatalln(fmt.Sprintf("Theme: \"%s\" was not found", _theme))
		}

		viper.Set("polybar.theme", _theme)
		Config.Polybar.Theme = _theme

		FullThemePath = FullThemesPath + "/" + _theme + "/config"
		main()

		// TODO: also check if theme is new and succeeded loading
		if themeIsValid {
			viper.WriteConfig()
		}

	},
}

func init() {
	polybarCmd.Flags().StringVarP(&_theme, "theme", "t", "", "Load a Polybar theme by name. The theme specified will be saved to dot's current_settings.")
	polybarCmd.Flags().BoolVarP(&_list, "list", "l", false, "Lists all themes found on the system.")
	polybarCmd.Flags().BoolVarP(&_select, "select", "s", false, "Select a theme interactively.")
	viper.BindPFlag("polybar.theme", polybarCmd.Flags().Lookup("theme"))
	// TODO: This is putting list on viper, which is then written to file
	// either figure out how to unmarshal Config and overwrite current settings with it,
	// or figure out how to reference flags without viper. 'list=true' doesn't belong in current_settings
	rootCmd.AddCommand(polybarCmd)
}

func listThemes() {
	for _, t := range InstalledPolybarThemes {
		fmt.Println(t)
	}
}

func findThemes() {
	if Config.Polybar.ThemesDirectory == "" {
		log.Fatalln("Please set polybar.themes_directory in current_settings.yml")
	}
	// TODO: validate theme in cs (current_settings.yml) exists on filesystem
	// TODO: if theme passed in exists and is different than the current theme, write it to cs

	// Look up installed themes.
	// A theme is considered to be installed if there is a directory with the themes name,
	// in the themes folder.
	f, err := ioutil.ReadDir(FullThemesPath)
	if err != nil {
		log.Errorln(err)
		log.Fatalf("Failed to read themes from %s", FullThemesPath)
	}

	for _, x := range f {
		if x.IsDir() && x.Name() != "global" {
			InstalledPolybarThemes = append(InstalledPolybarThemes, x.Name())
		}
	}
}

func validateTheme() bool {
	for _, x := range InstalledPolybarThemes {
		if _theme == x {
			return true
		}
	}
	return false
}

func main() {
	// connect to X server
	X, _ := xgb.NewConn()
	err := randr.Init(X)
	if err != nil {
		log.Fatal(err)
	}

	// get root node
	root := xproto.Setup(X).DefaultScreen(X).Root
	// get the resources of the screen
	resources, err := randr.GetScreenResources(X, root).Reply()
	if err != nil {
		log.Fatal(err)
	}
	// get the primary output
	primaryOutput, _ := randr.GetOutputPrimary(X, root).Reply()

	var displays []display
	// go through the connected outputs and get their position and resolution
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

	// order the displays by their position, left to right.
	sort.Slice(displays, func(i, j int) bool {
		return displays[i].xposition < displays[j].xposition
	})

	// kill all polybar sessions polybar
	cmd := exec.Command("sh", "-c", "killall -q polybar")
	out, err := cmd.CombinedOutput()
	if err != nil {
		// TODO: handle case where polybar isn't running yet
		// log.Fatalln("Failed to kill polybar")
		log.Println("Failed to kill polybar")
	}
	fmt.Println(string(out))

	// create the env vars we'll hand to polybar
	// polybar needs to know the theme, and what the left, right and main monitor are
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
	// add the theme to the environment
	t := fmt.Sprintf("polybar_theme=%s", FullThemePath)
	log.Infoln(t)
	polybarEnvVars = append(polybarEnvVars, t)

	// create a new array of env vars, appending the current environment
	// with the env vars created above
	newEnv := append(os.Environ(), polybarEnvVars...)

	// get the theme object for current theme from current_settings
	var theme Theme
	// TODO: maybe switch Themes to a map so i don't have to loop
	for _, t := range Config.Polybar.Themes {
		if Config.Polybar.Theme == t.Name {
			theme = t
		}
	}
	// Exit if it fails to find the theme object
	if theme.Name == "" {
		log.Error("Theme not found, exiting.")
		os.Exit(1)
	}

	// load bars from theme's .rasi file if none were specified in current_settings.yml
	var bars []string
	if len(theme.Bars) == 0 {
		log.Infoln("No bars specified in current-settings file. Auto-detecting bars...")
		bars = getBars(theme, FullThemePath)
	} else {
		log.Infoln("Bars specified in current-settings file...")
		bars = theme.Bars
	}
	// start all the bars
	var wg sync.WaitGroup
	for _, bar := range bars {
		log.Infoln(fmt.Sprintf("Loading bar '%s'", bar))
		wg.Add(1)
		go func(b string) {
			defer wg.Done()
			polybar(newEnv, b)
		}(bar)
	}
	wg.Wait()

	// g := getDefaultI3Gaps()
	adjustI3Gaps(theme.Gaps)
}

// Polybar themes can specify the gaps between i3 and the bar(s). This is useful
// when i3 doesn't respect the height of the bar, which happens when certain settings
// are enabled in polybar.
func adjustI3Gaps(g I3Gaps) {
	sides := []string{"top", "bottom", "left", "right"}
	sizes := []string{g.Top, g.Bottom, g.Left, g.Right}
	d := Config.I3wm.DefaultGaps
	defaults := []string{d.Top, d.Bottom, d.Left, d.Right}

	for i, s := range sides {
		var err error
		if sizes[i] == "" {
			log.Info(fmt.Sprintf("Setting i3wm \"%s\" gap to default: \"%s\"", s, defaults[i]))
			_, err = i3.RunCommand(fmt.Sprintf("gaps %s all set %s", s, defaults[i]))
		} else {
			log.Info(fmt.Sprintf("Setting i3wm \"%s\" gap to \"%s\", specified in theme: \"%s\"", s, sizes[i], _theme))
			_, err = i3.RunCommand(fmt.Sprintf("gaps %s all set %s", s, sizes[i]))
		}
		if err != nil {
			log.Errorln(err)
		}
	}
}

// func getDefaultI3Gaps() I3Gaps {
// 	f, err := os.Open(Home + "/" + Config.I3SettingsFile)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer f.Close()
// 	var g I3Gaps
// 	scanner := bufio.NewScanner(f)
// 	// re := regexp.MustCompile(`^gaps top`)
// 	for scanner.Scan() {
// 		// log.Info(scanner.Text())
// 		if strings.Contains(scanner.Text(), "gaps top") {
// 			log.Infoln(scanner.Text())
// 			// m := re.FindSubmatch(scanner.Bytes())
// 			// log.Infoln(string(m[0]))
// 			// b = append(b, string(m[1]))
// 		}
// 	}
// 	if err := scanner.Err(); err != nil {
// 		log.Errorln(err)
// 	}
// 	return g
// }

func getBars(theme Theme, path string) []string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var b []string
	scanner := bufio.NewScanner(f)
	// example: [bar/SOME.BAR] -> SOME.BAR
	re := regexp.MustCompile(`^\[bar\/(.*?)\]`)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "[bar/") {
			m := re.FindSubmatch(scanner.Bytes())
			b = append(b, string(m[1]))
		}
	}
	if len(b) == 0 {
		// TODO use a public variable to reference the path to current_settings.yml
		log.Fatalf("No bars found in:\n - %s\n - %s", FullThemePath, Home+"/code/dot/current_settings.yml")
	}
	return b
}

func polybar(env []string, bar string) {

	s := fmt.Sprintf("polybar -r %s", bar)
	cmd := exec.Command("bash", "-c", s)
	cmd.Env = env

	var stdoutBuf, stderrBuf bytes.Buffer
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()

	var errStdout, errStderr error
	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)
	err := cmd.Start()
	if err != nil {
		log.Errorf("cmd.Start() failed with '%s'\n", err)
	}

	_, errStdout = io.Copy(stdout, stdoutIn)
	_, errStderr = io.Copy(stderr, stderrIn)

	err = cmd.Wait()
	if err != nil {
		log.Errorf("Starting bar %s failed with %s\n", bar, err)
	}
	if errStdout != nil || errStderr != nil {
		log.Error("failed to capture stdout or stderr\n")
	}
	outStr, errStr := string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())
	fmt.Printf("stdout:\n%s\nerr:\n%s\n", outStr, errStr)
}
