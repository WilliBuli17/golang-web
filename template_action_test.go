package golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateActionIf(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFS(templates2, "templates/*.gohtml"))
	err := t.ExecuteTemplate(w, "if.gohtml", Page{
		Title: "Template Action If",
		Name:  "Willi",
	})
	if err != nil {
		panic(err)
	}
}

func TemplateActionRange(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFS(templates2, "templates/*.gohtml"))
	err := t.ExecuteTemplate(w, "range.gohtml", map[string]interface{}{
		"Title": "Template Action Range",
		"Hobbies": []string{
			"Game", "Read", "Code",
		},
	})
	if err != nil {
		panic(err)
	}
}

// jika ini dipakai untuk with maka sedikit dulit, karena templatenya tetap membaca alamat ini tetap ada
//func TemplateDataWith(w http.ResponseWriter, r *http.Request) {
//	t := template.Must(template.ParseFS(templates2, "templates/*.gohtml"))
//	err := t.ExecuteTemplate(w, "with.gohtml", Page{
//		Title: "Template Data Struct",
//		Name:  "Styephen William Buli",
//		Address: Address{
//			Street: "Jalanin dulu aja",
//			City:   "Fatamorgana",
//		},
//	})
//	if err != nil {
//		panic(err)
//	}
//}

func TemplateDataWith(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFS(templates2, "templates/*.gohtml"))
	err := t.ExecuteTemplate(w, "with.gohtml", map[string]interface{}{
		"Title": "Template Data Map",
		"Name":  "Willi Buli",
		"Address": map[string]interface{}{
			"Street": "Jalanin dulu aja",
			"City":   "Fatamorgana",
		},
	})
	if err != nil {
		panic(err)
	}
}

func TestTemplateAction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000", nil)
	recorder := httptest.NewRecorder()

	//TemplateActionIf(recorder, request) // untuk if
	//TemplateActionRange(recorder, request) // untuk range
	TemplateDataWith(recorder, request) // untuk with

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
