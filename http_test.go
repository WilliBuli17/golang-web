package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	_, err := fmt.Fprint(writer, "Hello World")
	if err != nil {
		panic(err)
	}
}

func TestHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/hello", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body)) // karena bodynya byte array jadi di konversi ke string

}
