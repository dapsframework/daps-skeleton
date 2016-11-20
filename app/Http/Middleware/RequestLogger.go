package Middleware

import (
	"net/http"
	"fmt"
)

type RequestLogger struct {
}

func (requestLogger *RequestLogger) Handle(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Beginning")
		handler.ServeHTTP(writer, request)
		fmt.Fprint(writer, "End")
	})
}
