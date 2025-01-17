package main

import (
	"github.com/itzngga/Roxy/core"
	_ "github.com/itzngga/Roxy/examples/cmd"
	"github.com/itzngga/Roxy/options"
	_ "github.com/mattn/go-sqlite3"
	"log"

	"os"
	"os/signal"
	"syscall"
)

func main() {
	app, err := core.NewGoRoxyBase(options.NewDefaultOptions())
	if err != nil {
		log.Fatal(err)
	}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	app.Shutdown()
}
