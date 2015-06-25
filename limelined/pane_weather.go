package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

var weatherSymbols = map[string]string{
	"tornado":               "☈",
	"tropical storm":        "☈",
	"hurricane":             "☈",
	"severe thunderstorms":  "☈",
	"thunderstorms":         "☈",
	"mixed rain and snow":   "☔",
	"mixed rain and sleet":  "☔",
	"mixed snow and sleet":  "",
	"freezing drizzle":      "☔",
	"drizzle":               "☔",
	"freezing rain":         "☔",
	"showers":               "☔",
	"snow flurries":         "☃",
	"light snow showers":    "☃",
	"blowing snow":          "☃",
	"snow":                  "☃",
	"hail":                  "☃",
	"sleet":                 "",
	"dust":                  "〰",
	"foggy":                 "〰",
	"haze":                  "〰",
	"smoky":                 "〰",
	"blustery":              "〰",
	"windy":                 "〰",
	"cold":                  "〇",
	"cloudy":                "☁",
	"mostly cloudy":         "☁",
	"mostly cloudy (night)": "☁",
	"mostly cloudy (day)":   "☁",
	"partly cloudy (night)": "☁",
	"partly cloudy (day)":   "☁",
	"clear (night)":         "☾",
	"sunny":                 "☼",
	"fair (night)":          "",
	"fair (day)":            "",
	"mixed rain and hail":   "☔",
	"hot": "☼",
	"isolated thunderstorms":  "☔",
	"scattered thunderstorms": "☈",
	"scattered showers":       "☔",
	"scattered snow showers":  "☃",
	"heavy snow":              "☃",
	"partly cloudy":           "☁",
	"thundershowers":          "☔",
	"snow showers":            "☃",
	"isolated thundershowers": "☈",
	"not available":           "",
}

func paneWeather() string {
	city_code, ok := readConfig("weather", "city_code").(string)
	if !ok {
		fmt.Println("error")
		return ""
	}
	res, err := http.Get(fmt.Sprintf("http://weather.yahooapis.com/forecastrss?w=%s&u=c", city_code))
	if err != nil {
		return ""
	}
	xml, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return ""
	}
	re := regexp.MustCompile("yweather:condition\\s*text=\"([^\"]+)\"\\s*code=\"([^\"]+)\"\\s*temp=\"([^\"]+)\"")
	matches := re.FindStringSubmatch(string(xml))
	if len(matches) == 4 {
		temp := matches[3]
		text := strings.ToLower(matches[1])
		weatherSymbol, ok := weatherSymbols[text]
		if !ok {
			weatherSymbol = ""
		}
		return fmt.Sprintf("%s %s", weatherSymbol, temp)
	}
	return ""
}
