package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")

	if name == "" {
		writer.WriteHeader(http.StatusBadRequest)
		_, err := fmt.Fprint(writer, "name is empty")
		if err != nil {
			panic(err)
		}
	} else {
		writer.WriteHeader(http.StatusOK)
		_, err := fmt.Fprintf(writer, "Hi, %s", name)
		if err != nil {
			panic(err)
		}
	}
}

func TestResponseCode(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/?name=willi", nil)
	recoeder := httptest.NewRecorder()

	ResponseCode(recoeder, request)

	response := recoeder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))
}
