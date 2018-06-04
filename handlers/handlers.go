package handlers

import (
	"net/http"
	"regexp"
	"html/template"
	"fmt"
	"io/ioutil"

	"github.com/adizim/webdemo/page"
	"github.com/adizim/webdemo/templates"
)

var validPath = regexp.MustCompile("^/(view|edit|save)/([A-Za-z0-9]+)$")

func MakeHandler(fn func(w http.ResponseWriter, r *http.Request, title string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func ViewHandler(w http.ResponseWriter, r *http.Request, title string) {	
	p, err := page.LoadPage(title)
	fmt.Printf("%v", title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	templates.RenderTemplate(w, "view", p)
}

func EditHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := page.LoadPage(title)
	if err != nil {
		p = &page.Page{Title: title}
	}
	templates.RenderTemplate(w, "edit", p)
}

func SaveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := []byte(r.FormValue("body"))
	p := &page.Page{Title: title, Body: body}
	err := p.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/" + title, http.StatusFound)
}

func titleToHtml(title []byte) []byte {
	var t string = string(title[:len(title) - len(".txt")])
	return []byte(fmt.Sprintf("<p><a href=\"/view/%s\">%s</a></p>", t, t))
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir("data")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var html []byte
	for _, file := range files {
		textFile := regexp.MustCompile("^([A-Za-z0-9]+).txt$")
		link := textFile.ReplaceAllFunc([]byte(file.Name()), titleToHtml)
		html = append(html, link...)
	}
	data := map[string]interface{} {
		"Html": template.HTML(string(html)),
	}
	templates.RenderTemplate(w, "index", data)
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/index", http.StatusFound)
}