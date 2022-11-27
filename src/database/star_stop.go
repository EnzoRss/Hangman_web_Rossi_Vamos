package database

import (
	"encoding/json"
	"fmt"
	"os"
)

func (w Data_Hangman) saveData() {
	f, _ := os.Create("save.txt")

	data := &Data_Hangman{
		Attempts:         w.Attempts,
		Word:             w.Word,
		ToFind:           w.ToFind,
		HangmanPositions: w.HangmanPositions,
	}
	save, _ := json.Marshal(data)
	f.WriteString(string(save))
	fmt.Print(string(save))
}

func (w *Data_Hangman) Resume() {
	var test Data_Hangman
	recup, _ := os.ReadFile(os.Args[2])
	json.Unmarshal([]byte(recup), &test)

	w.Attempts = test.Attempts
	w.ToFind = test.ToFind
	w.Word = test.Word
	w.HangmanPositions = test.HangmanPositions

	fmt.Println("Bon retour il vous reste :", 10-w.Attempts, "essai")
	DisplayArr(w.Word)
}
