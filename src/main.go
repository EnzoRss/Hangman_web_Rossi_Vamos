package main

import (
	"html/template"
	"net/http"
	"src/database"
)

func main() {
	var test database.Data_Hangman

	test.Init()

	tmpl_index := template.Must(template.ParseFiles("database/index.html"))
	tmpl_hangman := template.Must(template.ParseFiles("database/hangman.html"))
	http.HandleFunc("/Hangman", func(w http.ResponseWriter, r *http.Request) {
		tmpl_hangman.Execute(w, test)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl_index.Execute(w, nil)
	})
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs)) 

	http.ListenAndServe(":80", nil)

}
