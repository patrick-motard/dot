package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
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

	var s Settings
	json.Unmarshal(raw, &s)
	fmt.Println(s.Monitors.Current)

	ss := []string{s.Monitors.Location, s.Monitors.Current}
	cf := strings.Join(ss, "/")
	// fmt.Println(strings.Join(ss, "/"))
	fmt.Println(reflect.TypeOf(cf))
	r, err := ioutil.ReadFile(cf)

	fmt.Println("r:", reflect.TypeOf(r))
	// rs := string(r)
	fmt.Println(ss)
	fmt.Println(string(r))
	// return c
}
