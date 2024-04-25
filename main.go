package main

import (
	"net/http"
	"regexp"
	"simple-webapp/config"
	"simple-webapp/logger"
	"simple-webapp/page/handlers"

	"golang.org/x/exp/slog"
)

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

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
	cfg := config.Load()
	log := logger.Setup(cfg.Env)

	http.HandleFunc("/read/", makeHandler(handlers.Read))
	http.HandleFunc("/edit/", makeHandler(handlers.Edit))
	http.HandleFunc("/save/", makeHandler(handlers.Save))

	log.Info("server starting", slog.String("port", cfg.Port))
	if err := http.ListenAndServe(cfg.Port, nil); err != nil {
		log.Error("cannot start server", slog.String("error", err.Error()))
	}
	log.Error("server down")
}
