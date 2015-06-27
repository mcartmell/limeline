package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

const SOCKET_LOCATION = "/tmp/limelined.sock"

type paneData struct {
	paneId  int
	content string
}

type hub struct {
	update  chan bool
	content chan string
}

var currentPaneData []string

var h = hub{
	update:  make(chan bool),
	content: make(chan string),
}

func listenForPaneData() {
	currentPaneData = make([]string, len(Panes))
	c := make(chan paneData)
	startPaneFetcher(c)
	for {
		select {
		case pd := <-c:
			currentPaneData[pd.paneId] = pd.content
			fmt.Println("got pane data", pd)
		case <-h.update:
			h.content <- getCurrentStatus()
		}
	}
}

func startDaemon() {
	os.Remove(SOCKET_LOCATION)
	ln, err := net.ListenUnix("unix", &net.UnixAddr{SOCKET_LOCATION, "unix"})
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go readDaemon(conn)
	}
}

func readDaemon(c net.Conn) {
	defer c.Close()
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		switch scanner.Text() {
		case "status right":
			h.update <- true
			status := <-h.content
			c.Write([]byte(status))
		}
	}
}

func getCurrentStatus() string {
	status := getSep(-1, DEFAULT_BG)
	for _, s := range currentPaneData {
		status += s
	}
	return status + "\n"
}

func startPaneFetcher(c chan paneData) {
	for i, _ := range Panes {
		go fetchPane(c, i)
	}
}

// repeatedly gets content for a given pane at the interval specified in the config
func fetchPane(c chan paneData, i int) {
	interval := PaneConfig[Panes[i]]["interval"].(int)
	timer := time.Tick(time.Duration(interval) * time.Second)
	// get initial content
	go submitPaneContent(c, i)
	// start timer for subsequent content
	for _ = range timer {
		go submitPaneContent(c, i)
	}
}

func submitPaneContent(c chan paneData, i int) {
	content := getPaneContent(i)
	c <- paneData{
		paneId:  i,
		content: content,
	}
}

func daemonWorker() {
	loadConfig()
	go listenForPaneData()
	startDaemon()
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		switch args[0] {
		case "test":
			loadConfig()
			for i, _ := range Panes {
				fmt.Print(getPaneContent(i))
			}
		}
	} else {
		daemonize()
	}
}
