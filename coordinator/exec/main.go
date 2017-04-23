package main

import (
	"fmt"

	"github.com/jazaret/go-distributed/coordinator"
)

var dc *coordinator.DatabaseConsumer
var wc *coordinator.WebappConsumer

func main() {
	ea := coordinator.NewEventAggregator()
	dc = coordinator.NewDatabaseConsumer(ea)
	wc = coordinator.NewWebappConsumer(ea)
	ql := coordinator.NewQueueListener(ea)
	go ql.ListenForNewSource()

	fmt.Println("listening for new source...")

	var a string
	fmt.Scanln(&a)
}
