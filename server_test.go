package golang_web

import (
	"net/http"
	"testing"
)

// ----------------------------------------------------------------------------------------------------------------------
// cara buat server di golang
func TestServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:3000",
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
