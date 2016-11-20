package main

import (
	"net/http"
	"fmt"
	"github.com/dapsframework/dapsskeleton/app/Http/Controllers"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello, %s", request.URL.Path[1:])
}

func main() {
	finalHandler := http.HandlerFunc(handler)

	controller := new(Controllers.WelcomeController)
	controller.Init()

	var handle http.Handler = finalHandler

	for _, middleware := range controller.GetMiddleware() {
		handle = middleware.Handle(handle)
	}

	http.Handle("/", handle)
	http.ListenAndServe(":3000", nil)
}