package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Settings struct {
	Monitors string `json:"monitors"`
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

func main() {

	pages := GetSettings()
	for _, p := range pages {
		fmt.Println(p.toString())
	}

	fmt.Println(toJSON(pages))
}

func GetSettings() []Settings {
	raw, err := ioutil.ReadFile("/home/han/.config/dotfiles/settings.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c []Settings
	json.Unmarshal(raw, &c)
	return c
}
