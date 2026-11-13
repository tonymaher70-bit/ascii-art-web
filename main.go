package main

import (
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Result string
	Banner string
	Text   string
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 page not found", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, PageData{
		Banner: "standard",
		Text:   "",
	})
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if banner != "standard" &&
		banner != "shadow" &&
		banner != "thinkertoy" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	result, err := GenerateASCII(text, banner)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := PageData{
		Result: result,
		Banner: banner,
		Text:   text,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	var err error
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(home))
	mux.Handle("/ascii-art", http.HandlerFunc(asciiArtHandler))

	err = http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
