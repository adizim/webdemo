package templates

import (
	"net/http"
	"html/template"
	"fmt"
)

var templates = template.Must(template.ParseFiles("tmpl/edit.html", "tmpl/view.html", "tmpl/index.html"))

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmpl + ".html", data)
	if err != nil {
		fmt.Println("heres the error!")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}