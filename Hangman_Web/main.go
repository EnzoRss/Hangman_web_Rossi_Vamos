package main

import (
	"Hangman_Web/database"
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	var test database.Data_Hangman
	test.Init()

	tmpl_index := template.Must(template.ParseFiles("templates/index.html"))
	tmpl_hangman := template.Must(template.ParseFiles("templates/hangman.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl_index.Execute(w, nil)
	})

	http.HandleFunc("/Hangman", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			value := r.FormValue("lettre")
			fmt.Println(value)
			test.Input(value)
		}
		tmpl_hangman.Execute(w, test)
	})
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))

	http.ListenAndServe(":8080", nil)

}
