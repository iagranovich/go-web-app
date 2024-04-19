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

func handlerRead(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/read/"):]
	p, _ := getPage(title)

	t, _ := template.ParseFiles("read.html")
	t.Execute(w, p)
}

func handlerEdit(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := getPage(title)
	if err != nil {
		p = &page{Title: title}
	}

	t, _ := template.ParseFiles("edit.html")
	t.Execute(w, p)
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
