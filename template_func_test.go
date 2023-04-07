package golang_web

import (
	"html/template"
	"net/http"
	"strings"
)

type MyPage struct {
	Title string
	Name  string
}

func (myPage MyPage) SayHello(name string) string {
	return "Hello " + name + ", My Name is " + myPage.Name
}

func TemplateFunction(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFS(templates2, "templates/*.gohtml"))
	err := t.ExecuteTemplate(w, "function", MyPage{
		Title: "Template Function",
		Name:  "Buli",
	})
	if err != nil {
		panic(err)
	}
}

func TemplateGlobalFunction(w http.ResponseWriter, r *http.Request) {
	t := template.New("function_global")
	t = t.Funcs(template.FuncMap{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})

	t = template.Must(t.ParseFS(templates2, "templates/*.gohtml"))

	err := t.ExecuteTemplate(w, "function_global", MyPage{
		Title: "Template Function Global",
		Name:  "Buli",
	})
	if err != nil {
		panic(err)
	}
}
