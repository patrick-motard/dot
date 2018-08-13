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
	"strings"
)

var s Settings

type Settings struct {
	Monitors Monitors `json:"monitors"`
}

type Monitors struct {
	Current  string `json:"current"`
	Location string `json:"location"`
}

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
	e := ioutil.WriteFile("/home/han/.config/dotfiles/settings.json", jsonData, 0644)
	check(e)

}

func (m *Monitors) SetCurrent(current string) {
	location := m.Location
	files, err := ioutil.ReadDir(location)
	check(err)
	for i := 0; i < len(files); i++ {
		if current == files[i].Name() {
			fmt.Println("found a match!")
			m.Current = current
			screenlayoutScript := exec.Command("/bin/sh", strings.Join([]string{m.Location, m.Current}, "/"))
			err := screenlayoutScript.Run()
			fmt.Println(strings.Join([]string{m.Location, m.Current}, "/"))
			check(err)
			return
		}
	}
	fmt.Println("file not found")
}

func GetSettings() Settings {
	raw, err := ioutil.ReadFile("/home/han/.config/dotfiles/settings.json")
	// raw, err := ioutil.ReadFile("/home/han/.config/dotfiles/settings.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	json.Unmarshal(raw, &s)

	s.Monitors.Location = fullPath(s.Monitors.Location)
	// fmt.Println(s.Monitors.Current)

	return s

	// ss := []string{s.Monitors.Location, s.Monitors.Current}
	// cf := strings.Join(ss, "/")
	// fmt.Println(strings.Join(ss, "/"))
	// fmt.Println(reflect.TypeOf(cf))
	// r, err := ioutil.ReadFile(cf)
	// cf = fullPath(cf)
	// fmt.Println(fullPath(cf))

	// var  (
	// 	cmdOut []byte
	// 	err2   error
	// )
	// cmdName := "/bin/bash"
	// cmdArgs := []string{cf}
	// // cmd := exec.Command("bash", "/home/han/hello_world")
	// // fmt.Println(err)
	// if cmdOut, err2 = exec.Command(cmdName, cmdArgs...).Output(); err2 != nil {
	// 	fmt.Fprintln(os.Stderr, "There was an error running the feh script: ", err2)
	// 	os.Exit(1)
	// }

	// out := string(cmdOut)
	// fmt.Println("OUTPUT: " + out)

	// fmt.Println("r:", reflect.TypeOf(r))
	// rs := string(r)
	// fmt.Println("cf:", cf)
	// fmt.Println(string(r))
	// return c

	// cmd := exec.Command("/bin/bash", cf)
	// var out bytes.Buffer
	// cmd.Stderr = &out
	// err3 := cmd.Run()
	// if err3 != nil {
	// 	log.Fatal(err3)
	// }
	// fmt.Println("Output std: ", out.String())
}

func fullPath(s string) string {
	if strings.HasPrefix(s, "~/") {
		s = "/home/" + os.Getenv("USER") + strings.TrimPrefix(s, "~")
	}
	return s
}
