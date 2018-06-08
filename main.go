// Copyright © 2018 Patrick Motard <motard19@gmail.com>

package main

import "github.com/patrick-motard/dot/cmd"
import "github.com/patrick-motard/dot/lib"

func main() {
	lib.GetSettings()
	cmd.Execute()
}
