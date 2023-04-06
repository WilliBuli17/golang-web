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

type Page struct {
	Title   string
	Name    string
	Address Address
}

type Address struct {
	Street string
	City   string
}

//go:embed templates/*.gohtml
var templates2 embed.FS

func TemplateDataMap(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFS(templates2, "templates/*.gohtml"))
	err := t.ExecuteTemplate(w, "name.gohtml", map[string]interface{}{
		"Title": "Template Data Map",
		"Name":  "Willi Buli",
		"Address": map[string]interface{}{
			"Street": "Jalanin dulu aja",
		},
	})
	if err != nil {
		panic(err)
	}
}

func TemplateDataStruct(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFS(templates2, "templates/*.gohtml"))
	err := t.ExecuteTemplate(w, "name.gohtml", Page{
		Title: "Template Data Struct",
		Name:  "Styephen William Buli",
		Address: Address{
			Street: "Jalanin dulu aja",
		},
	})
	if err != nil {
		panic(err)
	}
}

func TestTemplateData(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000", nil)
	recorder := httptest.NewRecorder()

	//TemplateDataMap(recorder, request)    // untuk map
	TemplateDataStruct(recorder, request) // untuk struct

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
