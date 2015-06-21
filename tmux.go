// Handles the drawing of the tmux status pane
package main

import (
	"fmt"
)

// Prints the right statusbar
func printStatusRight() {
	var fgColour, bgColour string
	// print initial separator
	printSep(-1, DEFAULT_BG)
	for i, p := range Panes {
		// don't put separator on final panel
		includeSep := !(i == len(Panes)-1)
		cb := PaneConfig[p]["callback"].(func() string)
		pOpts, _ := PaneConfig[p]["options"].(paneOpts)

		// get foreground colour
		fgColour = pOpts.fgColour()

		// get background colour
		bgColour = pOpts.bgColour()

		content := cb()
		printRightPane(content, fgColour, bgColour)

		if includeSep {
			printSep(i, bgColour)
		}
	}
	return
}

func printSep(iPane int, bgColour string) {
	nextOpts, _ := PaneConfig[Panes[iPane+1]]["options"].(paneOpts)
	nextBG := nextOpts.bgColour()
	if nextBG == bgColour {
		// same background colour, use regular separator
		fmt.Print(rightSep(nextOpts.fgColour(), "", ""))
	} else {
		// diff background colour, use thick separator
		fmt.Print(rightSep(nextOpts.bgColour(), bgColour, ""))
	}
}

// Prints an individual pane
func printRightPane(content string, fgColour string, bgColour string) {
	fmt.Print(coloured(" "+content+" ", fgColour, bgColour))
}

// Helpers to format the status bar
func rightSep(fg string, bg string, char string) string {
	if bg == "" {
		return fmt.Sprintf("#[fg=%s]%s", fg, char)
	} else {
		return fmt.Sprintf("#[fg=%s,bg=%s]%s", fg, bg, char)
	}
}

func resetColour(str string) string {
	return ("#[default]" + str)
}

// colour a string, tmux style
func coloured(str string, fg string, bg string) string {
	return fmt.Sprintf("#[fg=%s,bg=%s]%s", fg, bg, str)
}
