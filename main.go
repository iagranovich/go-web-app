package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
)

var templates = template.Must(template.ParseFiles("edit.html", "read.html"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

type page struct {
	Title string
	Data  []byte
}

func (p *page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Data, 0600)
}

func getPage(title string) (*page, error) {
	file := title + ".txt"
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return &page{
		Title: title,
		Data:  data,
	}, nil
}

func loadTemplate(name string, w http.ResponseWriter, p *page) {
	err := templates.ExecuteTemplate(w, name, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handlerRead(w http.ResponseWriter, r *http.Request, title string) {
	p, err := getPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	loadTemplate("read.html", w, p)
}

func handlerEdit(w http.ResponseWriter, r *http.Request, title string) {
	p, err := getPage(title)
	if err != nil {
		p = &page{Title: title}
	}
	loadTemplate("edit.html", w, p)
}

func handlerSave(w http.ResponseWriter, r *http.Request, title string) {
	data := r.FormValue("data")
	p := &page{Title: title, Data: []byte(data)}
	if err := p.save(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/read/"+title, http.StatusFound)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func main() {
	http.HandleFunc("/read/", makeHandler(handlerRead))
	http.HandleFunc("/edit/", makeHandler(handlerEdit))
	http.HandleFunc("/save/", makeHandler(handlerSave))
	log.Fatal(http.ListenAndServe(":8089", nil))
}
