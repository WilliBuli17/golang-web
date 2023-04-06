package golang_web

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SimpleHTML(writer http.ResponseWriter, request *http.Request) {
	templateText := `<html><body>{{.}}</body></html>`
	//t, e := template.New("SIMPLE").Parse(templateText)
	//if e != nil {
	//	panic(e)
	//}
	//Cara simple dari yang di coment di atas
	t := template.Must(template.New("SIMPLE").Parse(templateText))

	err := t.ExecuteTemplate(writer, "SIMPLE", "Hello HTML Template")
	if err != nil {
		panic(err)
	}
}

func TestTemplate(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000", nil)
	recorder := httptest.NewRecorder()

	SimpleHTML(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

func SimpleHTMLFile(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/simple.gohtml")
	if err != nil {
		panic(err)
	}

	err = t.ExecuteTemplate(w, "simple.gohtml", "Hello HTML Template File")
	if err != nil {
		panic(err)
	}
}

func TestTemplateHTMLFile(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLFile(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

func SimpleHTMLDirectory(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseGlob("./templates/*.gohtml")
	if err != nil {
		panic(err)
	}

	err = t.ExecuteTemplate(w, "simple.gohtml", "Hello HTML Template Directory")
	if err != nil {
		panic(err)
	}
}

func TestTemplateHTMLDirectory(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLDirectory(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

//go:embed templates/*.gohtml
var templates embed.FS

func SimpleHTMLGoolangEmbed(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}

	err = t.ExecuteTemplate(w, "simple.gohtml", "Hello HTML Template Golang Embed")
	if err != nil {
		panic(err)
	}
}

func TestTemplateHTMLGolangEmbed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLGoolangEmbed(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
