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

func getPage(title string) *page {
	file := title + ".txt"
	data, _ := os.ReadFile(file)
	return &page{
		Title: title,
		Data:  data,
	}
}

func handlerRead(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/read/"):]
	p := getPage(title)
	t, _ := template.ParseFiles("template.html")
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/read/", handlerRead)
	log.Fatal(http.ListenAndServe(":8089", nil))
}
