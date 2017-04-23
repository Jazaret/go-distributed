package main

import (
	"fmt"

	"github.com/jazaret/go-distributed/coordinator"
)

var dc *coordinator.DatabaseConsumer

func main() {
	ea := coordinator.NewEventAggregator()
	dc := coordinator.NewDatabaseConsumer(ea)
	ql := coordinator.NewQueueListener(ea)
	go ql.ListenForNewSource()

	fmt.Println("listening for new source...")

	var a string
	fmt.Scanln(&a)
}
