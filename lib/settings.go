package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
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
	// r, err := ioutil.ReadFile(cf)
	cf = fullPath(cf)
	fmt.Println(fullPath(cf))

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
	fmt.Println("cf:", cf)
	// fmt.Println(string(r))
	// return c

	cmd := exec.Command("/bin/bash", cf)
	var out bytes.Buffer
	cmd.Stderr = &out
	err3 := cmd.Run()
	if err3 != nil {
		log.Fatal(err3)
	}
	fmt.Println("Output std: ", out.String())
}

func fullPath(s string) string {
	if strings.HasPrefix(s, "~/") {
		s = "/home/" + os.Getenv("USER") + strings.TrimPrefix(s, "~")
	}
	return s
}
