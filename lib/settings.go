package lib

import (
	// "bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	// "log"
	"os"
	"os/exec"
	// "reflect"
	log "github.com/sirupsen/logrus"
	"strings"
)

var s Settings
var settingsFilePath = "/home/han/code/go/src/github.com/patrick-motard/dot/current_settings.json"
var defaultSettingsFilePath = "/home/han/code/go/src/github.com/patrick-motard/dot/settings.json"

type Settings struct {
	Monitors Monitors `json:"monitors"`
	Sound    Sound    `json:"sound"`
}

type Monitors struct {
	Current  string `json:"current"`
	Location string `json:"location"`
}

type Sound struct {
	Port string `json:"port"`
}

// type Sound struct {
// 	DefaultSink Sink `json:"defaultSink"`
// }

// type Sink struct {
// 	Id int `json:"id"`
// 	CurrentPort string `json:"currentPort"`
// 	Ports []string `json:"ports"`
// }

func check(e error) {
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(1)
	}
}

func (p Settings) toString() string {
	return toJSON(p)
}

func toJSON(p interface{}) string {
	bytes, err := json.Marshal(p)
	check(err)
	return string(bytes)
}

func (s *Settings) WriteSettings() {
	jsonData, err := json.MarshalIndent(s, "", "    ")
	check(err)
	e := ioutil.WriteFile(settingsFilePath, jsonData, 0644)
	check(e)
}

func (m *Monitors) RunCurrent() {
	fullPath := strings.Join([]string{m.Location, m.Current}, "/")
	screenlayoutScript := exec.Command("/bin/sh", fullPath)
	e := screenlayoutScript.Run()
	check(e)
	return
}

func (m *Monitors) SetCurrent(current string) {
	location := m.Location
	// files, err := ioutil.ReadDir(location)
	files, err := ioutil.ReadDir(location)
	if err != nil {
		fmt.Printf(fmt.Sprintf("Error: Directory not found %s", location))
		os.Exit(1)
	}
	for i := 0; i < len(files); i++ {
		if current == files[i].Name() {
			fmt.Println("found a match!")
			m.Current = current
			return
		}
	}
	fmt.Println("file not found")
}

func (s *Settings) PrettyPrint() {
	j, err := json.MarshalIndent(s, "", "    ")
	check(err)
	fmt.Printf("%s\n", j)
}

func GetSettings() Settings {
	// TODO: replace these prints with debug/info level logs via glog or loggo, probbably glog for simplicity
	fmt.Println(fmt.Sprintf("Loading config: %s", settingsFilePath))
	raw, err := ioutil.ReadFile(settingsFilePath)

	if err != nil {
		log.Info("Failed to load current_settings.json.")
		// fmt.Println("Failed to load current_settings.json.")
		fmt.Println(fmt.Sprintf("Loading config: %s", defaultSettingsFilePath))
		raw, err = ioutil.ReadFile(defaultSettingsFilePath)
	}

	if err != nil {
		fmt.Println("Failed to read settings from settings.json & current_settings.json")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	err = json.Unmarshal(raw, &s)
	check(err)

	s.Monitors.Location = fullPath(s.Monitors.Location)

	return s

}

func fullPath(s string) string {
	if strings.HasPrefix(s, "~/") {
		s = "/home/" + os.Getenv("USER") + strings.TrimPrefix(s, "~")
	}
	return s
}
