package main

import (
	"html/template"
	"net/http"
)

type PageData struct {
	Result string
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

	tmpl.Execute(w, nil)
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
	}

	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/ascii-art", asciiArtHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
