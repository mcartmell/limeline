// Handles the drawing of the tmux status pane
package main

import (
	"fmt"
)

const DEFAULT_FG = "colour154"
const DEFAULT_BG = "colour234"

func toggleColour(i int, sw int) string {
	if i%2 == sw {
		return DEFAULT_BG
	} else {
		return DEFAULT_FG
	}
}

func getPaneContent(i int) string {
	var fgColour, bgColour string

	includeSep := !(i == len(Panes)-1)
	p := Panes[i]
	cb, ok := PaneCallbacks[p]
	if !ok {
		return ""
	}
	pOpts, _ := PaneConfig[p]["options"].(paneOpts)

	// get foreground colour
	if fgColour, ok = pOpts.fgColour(); !ok {
		fgColour = toggleColour(i, 1)
	}

	// get background colour
	if bgColour, ok = pOpts.bgColour(); !ok {
		bgColour = toggleColour(i, 0)
	}

	content := cb()

	s := getRightPane(content, fgColour, bgColour)

	if includeSep {
		s += getSep(i, bgColour)
	}
	return s
}

func getSep(iPane int, bgColour string) (s string) {
	nextOpts, _ := PaneConfig[Panes[iPane+1]]["options"].(paneOpts)
	nextBG, ok := nextOpts.bgColour()
	if !ok {
		nextBG = toggleColour(iPane+1, 0)
	}
	nextFG, ok := nextOpts.fgColour()
	if !ok {
		nextFG = toggleColour(iPane+1, 1)
	}
	if nextBG == bgColour {
		// same background colour, use regular separator
		s = rightSep(nextFG, "", "")
	} else {
		// diff background colour, use thick separator
		s = rightSep(nextBG, bgColour, "")
	}
	return s
}

// Gets an individual pane
func getRightPane(content string, fgColour string, bgColour string) string {
	return coloured(" "+content+" ", fgColour, bgColour)
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
