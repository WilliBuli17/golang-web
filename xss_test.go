package golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateAutoEscape(w http.ResponseWriter, r *http.Request) {
	onceInitTemplate.Do(func() {
		templateMutex.Lock()
		defer templateMutex.Unlock()
		myTemplates = initTemplates()
	})

	err := myTemplates.ExecuteTemplate(w, "xss", map[string]interface{}{
		"Title": "Template XSS",
		"Body":  "<p>Ini Body<script>alert('Anda di Heck')</script></p>",
	})

	if err != nil {
		panic(err)
	}
}

func TemplateAutoEscapeDisabled(w http.ResponseWriter, r *http.Request) {
	onceInitTemplate.Do(func() {
		templateMutex.Lock()
		defer templateMutex.Unlock()
		myTemplates = initTemplates()
	})

	err := myTemplates.ExecuteTemplate(w, "xss", map[string]interface{}{
		"Title": "Template XSS",
		"Body":  template.HTML("<p>Ini Body</p>"),
	})

	if err != nil {
		panic(err)
	}
}

func TemplateXSS(w http.ResponseWriter, r *http.Request) {
	onceInitTemplate.Do(func() {
		templateMutex.Lock()
		defer templateMutex.Unlock()
		myTemplates = initTemplates()
	})

	err := myTemplates.ExecuteTemplate(w, "xss", map[string]interface{}{
		"Title": "Template XSS",
		"Body":  template.HTML(r.URL.Query().Get("body")),
	})

	if err != nil {
		panic(err)
	}
}

func TestTemplateAutoEscape(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/?body=<p>alert</p>", nil)
	recorder := httptest.NewRecorder()

	//TemplateAutoEscape(recorder, request) // ini yang auto escape
	//TemplateAutoEscapeDisabled(recorder, request) // ini jika ingin auto escape mati
	TemplateXSS(recorder, request) // ini contoh kasus xss

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

func TestTemplateAutoEscapeServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:3000",
		//Handler: http.HandlerFunc(TemplateAutoEscape),
		//Handler: http.HandlerFunc(TemplateAutoEscapeDisabled),
		Handler: http.HandlerFunc(TemplateXSS),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
