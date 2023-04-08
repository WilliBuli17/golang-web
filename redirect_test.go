package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func RedirecttTo(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello Redirect")
	if err != nil {
		panic(err)
	}
}

func RedirectFrom(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/redirect-to", http.StatusTemporaryRedirect)
}

func TestRedirect(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/redirect-from", RedirectFrom)
	mux.HandleFunc("/redirect-to", RedirecttTo)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
