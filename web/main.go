package main

import (
	"net/http"

	"github.com/jazaret/go-distributed/web/controller"
)

func main() {
	controller.Initialize()

	http.ListenAndServe(":3000", nil)
}
