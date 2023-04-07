package golang_web

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
)

//go:embed templates/*.gohtml
var templatesFile embed.FS

var myTemplates *template.Template
var templateMutex sync.Mutex
var onceInitTemplate sync.Once

type PageTemplate struct {
	Title string
	Name  string
}

func (PageTemplate PageTemplate) SayHello(name string) string {
	return "Hello " + name + ", My Name is " + PageTemplate.Name
}

func initTemplates() *template.Template {
	t := template.New("")
	t = t.Funcs(template.FuncMap{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})

	return template.Must(t.ParseFS(templatesFile, "templates/*.gohtml"))
}

func TemplateFunctionCaching(w http.ResponseWriter, r *http.Request) {
	onceInitTemplate.Do(func() {
		templateMutex.Lock()
		defer templateMutex.Unlock()
		myTemplates = initTemplates()
	})

	err := myTemplates.ExecuteTemplate(w, "function", PageTemplate{
		Title: "Template Function",
		Name:  "Buli",
	})
	if err != nil {
		panic(err)
	}
}

func TemplateGlobalFunctionCaching(w http.ResponseWriter, r *http.Request) {
	onceInitTemplate.Do(func() {
		templateMutex.Lock()
		defer templateMutex.Unlock()
		myTemplates = initTemplates()
	})

	err := myTemplates.ExecuteTemplate(w, "function_global", PageTemplate{
		Title: "Template Function Global",
		Name:  "Buli",
	})
	if err != nil {
		panic(err)
	}
}

func TestTemplateFunctionCaching(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionCaching(recorder, request)
	//TemplateGlobalFunctionCaching(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
