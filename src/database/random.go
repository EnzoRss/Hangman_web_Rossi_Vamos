package database

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type Data_Hangman struct {
	Word             []string   // Word composed of '_', ex: H_ll_
	ToFind           string     // Final word chosen by the program at the beginning. It is the word to find
	Attempts         int        // Number of attempts left
	HangmanPositions [11]string // It can be the array where the positions parsed in "hangman.txt" are stored
}

func (w *Data_Hangman) ChoseWord() {
	arr, _ := os.ReadFile("database/Word.txt")
	arrWord := strings.Split(string(arr), string(13))
	w.ToFind = arrWord[rand.Intn(len(arrWord))]
}

func (w *Data_Hangman) DisplayLetters() {
	n := len(w.ToFind)/2 - 1
	for i := 0; i < len(w.ToFind)-1; i++ {
		w.Word = append(w.Word, "_")
	}
	for i := 0; i < n; i++ {
		nb := rand.Intn(len(w.ToFind))
		w.Word[nb-1] = string(w.ToFind[nb])
	}
	fmt.Print("Voici votre mots avec les lettre donner :")
	DisplayArr(w.Word)
}

func DisplayArr(arr []string) {
	for _, str := range arr {
		fmt.Print(str)
	}
	fmt.Print("\n")
}

func (w *Data_Hangman) Init() {
	w.ChoseWord()
	w.DisplayLetters()
	w.Position_init()
}
