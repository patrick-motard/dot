package cmd

import (
	"sort"

	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/randr"
	"github.com/BurntSushi/xgb/xproto"
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
type displays struct {
	displays []display
	left     display
	right    display
	primary  display
}

func (ds displays) getLeft() display {
	ds.get()
	return ds.left
}
func (ds displays) getRight() display {
	ds.get()
	return ds.right
}
func (ds displays) getPrimary() display {
	ds.get()
	return ds.primary
}

// gets current displays, sets them on the struct, and returns them
func (ds *displays) get() []display {
	if len(ds.displays) > 0 {
		return ds.displays
	}
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

			ds.displays = append(ds.displays, d)
		}
	}
	// order the displays by their position, left to right.
	sort.Slice(ds.displays, func(i, j int) bool {
		return ds.displays[i].xposition < ds.displays[j].xposition
	})
	for i, d := range ds.get() {
		// skip inactive monitors
		if !d.active {
			continue
		}
		if d.primary {
			ds.primary = d
		} else if i == 0 {
			ds.left = d
		} else if i == 1 || i == 2 {
			ds.right = d
		}
	}
	return ds.displays
}
