package main

import (
	"html/template"
	"net/http"
	// "src/database"
)

func main() {
	//var test database.Data_Hangman

	tmpl_index := template.Must(template.ParseFiles("database/index.html"))
	tmpl_hangman := template.Must(template.ParseFiles("database/hangman.html"))
	http.HandleFunc("/Hangman", func(w http.ResponseWriter, r *http.Request) {
		tmpl_hangman.Execute(w, nil)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl_index.Execute(w, nil)
	})

	http.ListenAndServe(":80", nil)

}
