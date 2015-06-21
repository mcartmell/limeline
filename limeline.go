// A minimalist statusbar tool for tmux, inspired by Powerline
package main

import (
	"flag"
	"log"
)

const DEFAULT_FG = "colour154"
const DEFAULT_BG = "colour234"

// Display these panes in order
var Panes = []string{"loadavg", "sghaze", "datetime"}

// Functions to call to load each pane
var PaneConfig = map[string]map[string]interface{}{
	"loadavg": {
		"callback": paneLoadAvg,
	},
	"sghaze": {
		"callback": paneSGHaze,
	},
	"datetime": {
		"callback": paneDateTime,
	},
}

// Main program. Only supports the right statusbar for now.
func main() {
	flag.Parse()
	loadConfig()
	args := flag.Args()
	if len(args) > 0 {
		switch args[0] {
		case "right":
			printStatusRight()
		default:
			log.Fatal("Usage: limeline right")
		}
	} else {
		log.Fatal("Usage: limeline right")
	}
}
