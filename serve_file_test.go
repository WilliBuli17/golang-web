package golang_web

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func ServeFile(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("name") != "" {
		http.ServeFile(writer, request, "./resources/ok.html")
	} else {
		http.ServeFile(writer, request, "./resources/notfound.html")
	}
}

func TestServeFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources/ok.html
var resourceOK string

//go:embed resources/notfound.html
var resourceNotFound string

func ServeFileEmbed(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("name") != "" {
		fmt.Fprintf(writer, resourceOK)
	} else {
		fmt.Fprintf(writer, resourceNotFound)
	}
}

func TestServeFileEmbed(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
