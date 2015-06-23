// A minimalist statusbar tool for tmux, inspired by Powerline
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
)

// Main program. Only supports the right statusbar for now.
func main() {
	flag.Parse()
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

func printStatusRight() {
	c, err := net.DialUnix("unix", nil, &net.UnixAddr{"/tmp/limelined.sock", "unix"})
	if err != nil {
		fmt.Println("error", err)
	} else {
		_, err = c.Write([]byte("status right\n"))
		if err != nil {
			fmt.Println("error", err)
		} else {
			reader := bufio.NewReader(c)
			if s, err := reader.ReadString('\n'); err == nil {
				fmt.Println(s)
			}
		}
	}
}
