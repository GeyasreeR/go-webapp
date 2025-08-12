package main

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func renderPage(w http.ResponseWriter, page string, title string) {
	tmpl, err := template.ParseFiles(
		filepath.Join("templates", "layout.html"),
		filepath.Join("templates", page+".html"),
	)
	if err != nil {
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, map[string]string{"Title": title})
}

// ✅ Named route handlers
func homeHandler(w http.ResponseWriter, r *http.Request) {
	renderPage(w, "home", "Home")
}

func courseHandler(w http.ResponseWriter, r *http.Request) {
	renderPage(w, "course", "Courses")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	renderPage(w, "about", "About")
}

func blogHandler(w http.ResponseWriter, r *http.Request) {
	renderPage(w, "blog", "Blog")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/course", courseHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/blog", blogHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	println("🚀 Server running at http://127.0.0.1:8080")
	http.ListenAndServe(":8080", nil)
}
