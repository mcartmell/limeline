package main

import (
	"github.com/mcartmell/go-daemon"
	"log"
)

func daemonize() {
	cntxt := &daemon.Context{
		PidFileName: "/tmp/limelined.pid",
		PidFilePerm: 0644,
		LogFileName: "",
		LogFilePerm: 0640,
		WorkDir:     "/tmp",
		Umask:       027,
		Args:        []string{"[limelined]"},
	}

	if len(daemon.ActiveFlags()) > 0 {
		d, err := cntxt.Search()
		if err != nil {
			log.Fatalln("Unable send signal to the daemon:", err)
		}
		daemon.SendCommands(d)
		return
	}

	d, err := cntxt.Reborn()
	if err != nil {
		log.Println(5, err)
		return
	}
	if d != nil {
		log.Println(4, err)
		return
	}
	defer cntxt.Release()

	log.Println("- - - - - - - - - - - - - - -")
	log.Println("daemon started")

	go daemonWorker()

	err = daemon.ServeSignals()
	if err != nil {
		log.Println("Error:", err)
	}
	log.Println("daemon terminated")
}
