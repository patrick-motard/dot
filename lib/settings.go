package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Settings struct {
	Monitors Monitors
}

type Monitors struct {
	Current  string `json:"current"`
	Location string `json:"location"`
}

func (p Settings) toString() string {
	return toJSON(p)
}

func toJSON(p interface{}) string {
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return string(bytes)
}

// func main() {

// 	settings := GetSettings()
// 	for _, p := range settings {
// 		fmt.Println(p.toString())
// 	}

// 	fmt.Println(toJSON(settings))
// }

func GetSettings() {
	raw, err := ioutil.ReadFile("/home/han/.config/dotfiles/settings.json")
	// raw, err := ioutil.ReadFile("/home/han/.config/dotfiles/settings.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c Settings
	json.Unmarshal(raw, &c)
	fmt.Println(c.Monitors.Current)
	// return c
}
