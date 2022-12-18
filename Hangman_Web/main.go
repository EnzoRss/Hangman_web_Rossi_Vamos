package main

import (
	"Hangman_Web/database"
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	var test =new(database.Data_Hangman)
	cmpt := 0

	tmpl_index := template.Must(template.ParseFiles("templates/index.html"))
	tmpl_hangman := template.Must(template.ParseFiles("templates/hangman.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl_index.Execute(w, nil)
	})

	http.HandleFunc("/Hangman",
		func(w http.ResponseWriter, r *http.Request) {
			if cmpt == 0 {
				test.Username = r.URL.Query().Get("username")
				test.Level = r.FormValue("level")
				test.Init()
				cmpt++
			}
			if r.Method == "POST" {
				value := r.FormValue("lettre")
				rejouer := r.FormValue("rejouer")
				fmt.Println("rejouer :",rejouer)
				if rejouer == "rejouer" && cmpt >0 {
					fmt.Println("ici")
				 	test.ReInit()
				 }  else{
					fmt.Println(value)
					test.Input(value)
				}
			}
			tmpl_hangman.Execute(w, test)
		})
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))

	http.ListenAndServe(":8080", nil)

}
