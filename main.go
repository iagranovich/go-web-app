package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

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

func loadTemplate(name string, w http.ResponseWriter, p *page) error {
	t, err := template.ParseFiles(name)
	t.Execute(w, p)
	return err
}

func handlerRead(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/read/"):]
	p, err := getPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	loadTemplate("read.html", w, p)
}

func handlerEdit(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := getPage(title)
	if err != nil {
		p = &page{Title: title}
	}
	loadTemplate("edit.html", w, p)
}

func handlerSave(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	data := r.FormValue("data")
	p := &page{Title: title, Data: []byte(data)}
	p.save()
	http.Redirect(w, r, "/read/"+title, http.StatusFound)
}

func main() {
	http.HandleFunc("/read/", handlerRead)
	http.HandleFunc("/edit/", handlerEdit)
	http.HandleFunc("/save/", handlerSave)
	log.Fatal(http.ListenAndServe(":8089", nil))
}
