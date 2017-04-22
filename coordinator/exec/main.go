package main

import (
	"fmt"

	"github.com/jazaret/go-distributed/coordinator"
)

func main() {
	ql := coordinator.NewQueueListener()
	go ql.ListenForNewSource()

	fmt.Println("listening for new source...")

	var a string
	fmt.Scanln(&a)
}
