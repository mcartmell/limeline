// Code to fetch pane data for the statusbar
package main

import (
	"encoding/json"
	"fmt"
	"github.com/cloudfoundry/gosigar" // for system info
	"io/ioutil"
	"net/http"
	"time"
)

// Display these panes in order
var Panes = []string{"loadavg", "sghaze", "datetime"}

// Functions to call to load each pane
var PaneConfig = map[string]map[string]interface{}{
	"loadavg": {
		"interval": 5,
	},
	"sghaze": {
		"interval": 60,
	},
	"datetime": {
		"interval": 2,
	},
	"weather": {
		"interval": 60,
	},
}

var PaneCallbacks = map[string]func() string{
	"loadavg":  paneLoadAvg,
	"sghaze":   paneSGHaze,
	"datetime": paneDateTime,
	"weather":  paneWeather,
}

// Simple load average pane
func paneLoadAvg() string {
	avg := sigar.LoadAverage{}
	avg.Get()
	return fmt.Sprintf("%.2f %.2f %.2f", avg.One, avg.Five, avg.Fifteen)
}

// Singapore haze. See https://github.com/mcartmell/powerline-segment-sghaze
func paneSGHaze() string {
	res, err := http.Get("http://sghaze.herokuapp.com")
	if err != nil {
		return ""
	}
	jsonb, err := ioutil.ReadAll(res.Body)
	defer func() {
		res.Body.Close()
	}()
	if err != nil {
		return ""
	}
	content := make(map[string]interface{})
	if err := json.Unmarshal(jsonb, &content); err != nil {
		return ""
	}
	if haze, ok := content["Central"]; ok {
		if hazeVal, ok := haze.(string); ok {
			return fmt.Sprintf("░ %s", hazeVal)
		}
	}
	return ""
}

// Simple datetime pane
func paneDateTime() string {
	return time.Now().Format("2006-01-02 15:04")
}
