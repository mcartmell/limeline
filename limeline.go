// A minimalist statusbar tool for tmux, inspired by Powerline
package main

import (
	"flag"
	"fmt"
)

// Display these panes in order
var Panes = [...]string{"loadavg", "sghaze", "datetime"}

// Functions to call to load each pane
var Callbacks = map[string]func() string{
	"loadavg":  paneLoadAvg,
	"sghaze":   paneSGHaze,
	"datetime": paneDateTime,
}

// Main program. Only supports the right statusbar for now.
func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		switch args[0] {
		case "right":
			printStatusRight()
		}
	}
}

// Prints the right statusbar
func printStatusRight() {
	for i, p := range Panes {
		includeSep := !(i == len(Panes)-1)
		content := Callbacks[p]()
		printRightPane(content, includeSep)
	}
	return
}

// Prints an individual pane
func printRightPane(content string, includeSep bool) {
	fmt.Print(resetColour(content))
	if includeSep {
		fmt.Print(rightSep())
	}
}

// Helpers to format the status bar
func rightSep() string {
	return "#[fg=colour244,bg=colour234]  "
}

func resetColour(str string) string {
	return ("#[default]" + str)
}
