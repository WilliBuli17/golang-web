package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "X-PZN-Name"                     // nama cookie
	cookie.Value = request.URL.Query().Get("name") // value yang di set ke cookie
	cookie.Path = "/"                              // dia bisa diakses/aktif dimana saja atau di semua url kalau di set /

	http.SetCookie(writer, cookie)
	_, err := fmt.Fprint(writer, "Success Create Cookie")
	if err != nil {
		panic(err)
	}
}

func GetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("X-PZN-Name")

	if err != nil {
		_, err := fmt.Fprint(writer, "No Cookie")
		if err != nil {
			panic(err)
		}
	} else {
		_, err := fmt.Fprintf(writer, "Hello %s", cookie.Value)
		if err != nil {
			panic(err)
		}
	}
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/?name=Willi", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	cookies := recorder.Result().Cookies()

	for _, cookie := range cookies {
		fmt.Printf("Cookie %s : %s\n", cookie.Name, cookie.Value)
	}
}

func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000", nil)
	recorder := httptest.NewRecorder()

	cookie := new(http.Cookie)
	cookie.Name = "X-PZN-Name"
	cookie.Value = "Willi Buli"
	cookie.Path = "/"

	request.AddCookie(cookie)

	GetCookie(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
