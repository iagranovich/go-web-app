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

func getPage(title string) (*page, error) {
	file := title + ".txt"
	data, _ := os.ReadFile(file)
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
		p = &page{Title: title, Data: []byte{}}
	}
	t, _ := template.ParseFiles("edit.html")
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/read/", handlerRead)
	http.HandleFunc("/edit/", handlerEdit)
	log.Fatal(http.ListenAndServe(":8089", nil))
}
