package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}

	firstName := request.PostForm.Get("first_name")
	lastName := request.PostForm.Get("last_name")

	//cara otomatis, biat lebih singkat dari yang diatas
	//firstName := request.PostFormValue("first_name")
	//lastName := request.PostFormValue("last_name")

	_, e := fmt.Fprintf(writer, "Hello %s %s", firstName, lastName)
	if err != nil {
		panic(e)
	}
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("first_name=Willi&last_name=Buli")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
