package Middleware

import (
	"net/http"
	"fmt"
)

type ExtraContent struct {

}

func (requestLogger *ExtraContent) Handle(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "EXTRA")
		handler.ServeHTTP(writer, request)
		fmt.Fprint(writer, "CONTENT")
	})
}

