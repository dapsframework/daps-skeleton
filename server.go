package main

import (
	"net/http"
	"fmt"
	//"github.com/dapsframework/dapsskeleton/app/Http/Controllers"
	"github.com/dapsframework/daps/Server"
	"github.com/dapsframework/daps/Foundation"
	"path"
	"path/filepath"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello, %s", request.URL.Path[1:])
}

func main() {
	filename := filepath.Dir("")
	directory, _ := filepath.Abs(filename)

	application := Foundation.NewApplication(path.Join(directory, "src", "github.com", "dapsframework", "dapsskeleton"))

	server := Server.Server{
		Pattern: "/",
		Uri: ":3000",
		Handler: http.HandlerFunc(handler),
		Application: application,
	}

	server.Serve()

	//finalHandler := http.HandlerFunc(handler)
	//
	//controller := new(Controllers.WelcomeController)
	//controller.Init()
	//
	//var handle http.Handler = finalHandler
	//
	//for _, middleware := range controller.GetMiddleware() {
	//	handle = middleware.Handle(handle)
	//}
}