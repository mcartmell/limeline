// Loads config from YAML
package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
)

type paneOpts struct {
	Fg *string
	Bg *string
}

type paneConfig map[string]map[string]interface{}

func (self *paneConfig) get(pane string, key string) (interface{}, bool) {
	_, ok := PaneConfig[pane]
	if ok {
		if iface, ok := PaneConfig[pane]["raw"]; ok {
			if cfg, ok := iface.(map[interface{}]interface{}); ok {
				if val, ok := cfg[key]; ok {
					return val, true
				}
			}
		}
	}
	return nil, false
}

func (self *paneConfig) getString(pane string, key string) (string, bool) {
	if val, ok := self.get(pane, key); ok {
		if str, ok := val.(string); ok {
			return str, true
		}
	}
	return "", false
}

var validPaneOpts = [...]string{"fg", "bg"}

func readConfig(pane string, key string) (interface{}, bool) {
	_, ok := PaneConfig[pane]
	if ok {
		if iface, ok := PaneConfig[pane]["raw"]; ok {
			if cfg, ok := iface.(map[interface{}]interface{}); ok {
				if val, ok := cfg[key]; ok {
					return val, true
				}
			}
		}
	}
	return nil, false
}

func loadConfig() {
	cfg_file := fmt.Sprintf("%s/.config/limeline/config.yaml", os.Getenv("HOME"))
	if _, err := os.Stat(cfg_file); err != nil {
		// File doesn't exist
		return
	}
	if cfg_contents, err := ioutil.ReadFile(cfg_file); err != nil {
		log.Fatal(err)
	} else {
		cfg := make(map[string]interface{})
		if err := yaml.Unmarshal(cfg_contents, cfg); err != nil {
			log.Fatal(err)
		}
		// Add config options to plugins if they exist
		if plugcfg, ok := cfg["plugins"].(map[interface{}]interface{}); ok {
			for key, _ := range PaneConfig {
				// get plugin-specific config
				if paneCfg, ok := plugcfg[key]; ok {
					// check that the config is also a map
					if paneCfg, ok := paneCfg.(map[interface{}]interface{}); ok {
						po := paneOpts{}
						st := reflect.ValueOf(&po).Elem()
						// fill struct with valid pane options
						for _, opt := range validPaneOpts {
							if v, ok := paneCfg[opt]; ok {
								if str, ok := v.(string); ok {
									stv := st.FieldByName(strings.ToUpper(string(opt[0])) + opt[1:])
									stv.Set(reflect.ValueOf(&str))
								}
							}
						}
						PaneConfig[key]["options"] = po
						PaneConfig[key]["raw"] = paneCfg
					}
				}
			}
		}
		// Get the list of panes that we should use
		if panes, ok := cfg["panes"].([]interface{}); ok {
			newPanes := make([]string, len(panes))
			for i, v := range panes {
				if newPane, ok := v.(string); ok {
					newPanes[i] = newPane
				}
			}
			Panes = newPanes
		}
	}
}

func (self *paneOpts) fgColour() (fgColour string) {
	if self.Fg == nil {
		return "", false
	} else {
		fgColour = *self.Fg
	}
	return fgColour, true
}

func (self *paneOpts) bgColour() (bgColour string) {
	if self.Bg == nil {
		return "", false
	} else {
		bgColour = *self.Bg
	}
	return bgColour, true
}
