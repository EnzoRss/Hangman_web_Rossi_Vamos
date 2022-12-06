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
		tmpl_hangman.Execute(w, test)
		if r.Method == "POST" {
			value := r.FormValue("lettre")
			fmt.Println(value)
			test.Input(value)
			fmt.Print(test.Word_Display)
			r.Form.Get(test.Word_Display)
		}
	})
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))

	http.ListenAndServe(":8080", nil)

}
