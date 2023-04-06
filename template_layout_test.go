package golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateLayout(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFS(templates2, "templates/*.gohtml"))
	err := t.ExecuteTemplate(w, "layout_child", map[string]interface{}{
		"Title": "Template Action Range",
		"Name":  "Styephen William Buli",
	})
	if err != nil {
		panic(err)
	}
}

func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000", nil)
	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
