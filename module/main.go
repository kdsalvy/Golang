package main

import (
	"fmt"
	"net/http"

	"github.com/kd/learn/go/basics"
	"github.com/kd/learn/go/collections"
	"github.com/kd/learn/go/controller"
	"github.com/kd/learn/go/loops"
)

func main() {
	basics.Demo()
	collections.Demo()
	loops.Loop()

	port := ":3000"
	isStarted, error := startWebServer(port, 2)
	fmt.Println("Server start", isStarted, ", errors(if any)", error)
}

func startWebServer(port string, numberOfRetries int) (bool, error) {
	fmt.Println("Starting the webserver...")

	controller.RegisterControllers()
	http.ListenAndServe(port, nil)

	fmt.Println("Started the webserver on port", port)
	fmt.Println("Number of retries", numberOfRetries)
	return true, nil
}
