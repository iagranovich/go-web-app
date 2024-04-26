package handlers

import (
	"html/template"
	"net/http"
	"simple-webapp/page/model"
	"simple-webapp/page/storage"
)

var templates = template.Must(
	template.ParseFiles(
		"templates/edit.html",
		"templates/read.html"))

func Read(w http.ResponseWriter, r *http.Request, title string) {
	p, err := storage.GetPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	loadTemplate("read.html", w, p)
}

func Edit(w http.ResponseWriter, _ *http.Request, title string) {
	p, err := storage.GetPage(title)
	if err != nil {
		p = &model.Page{Title: title}
	}
	loadTemplate("edit.html", w, p)
}

func Save(w http.ResponseWriter, r *http.Request, title string) {
	data := r.FormValue("data")
	p := &model.Page{Title: title, Data: []byte(data)}
	if err := storage.SavePage(p); err != nil {
		//TO DO: add log
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/read/"+title, http.StatusFound)
}

func loadTemplate(name string, w http.ResponseWriter, p *model.Page) {
	err := templates.ExecuteTemplate(w, name, p)
	if err != nil {
		//TO DO: add log
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
