package main

import (
	"net/http"
	"webfrmw/html"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// auth users middleware? (this is just an example)
	r.Use((&html.Middlewares{}).Auth)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		html.New(w, r).With("home.html", nil).Render("index.html")
	})
	r.Get("/about", func(w http.ResponseWriter, r *http.Request) {
		vd := map[string]interface{}{}
		vd["component1data"] = map[string]interface{}{
			"component1_name":  "About",
			"component1_value": "Value",
		}
		html.New(w, r).With("about.html", vd).Render("index.html")
	})
	r.Get("/part", func(w http.ResponseWriter, r *http.Request) {
		html.New(w, r).With("home.html", nil).Render("item-part")
	})
	http.ListenAndServe(":3000", r)
}
