package database

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Data_Hangman struct {
	Word             []string   // Word composed of '_', ex: H_ll_
	ToFind           string     // Final word chosen by the program at the beginning. It is the word to find
	Attempts         int        // Number of attempts left
	HangmanPositions [11]string // It can be the array where the positions parsed in "hangman.txt" are stored
	Word_Display     string     // The word which use for the display
	Propo_let        []string   // list of the letter already propose
	Level            string     // level of the game
	Username         string     // username of the player
	path_file        string     //path of the file use for the game
	Result			 string     // result of the input enter on the templates
	Win				 int		// number of win
	Lose			 int 		// number of lose
	Point	 		 int 		// number of point 
	NbAttempts       int
}

func (w *Data_Hangman) ChoseWord() {
	rand.Seed(time.Now().UnixNano())
	arr, _ := os.ReadFile(w.path_file)
	arrWord := strings.Split(string(arr), "\n")
	w.ToFind = arrWord[rand.Intn(len(arrWord))]
	fmt.Println(w.ToFind)
}

func (w *Data_Hangman) DisplayLetters() {
	rand.Seed(time.Now().UnixNano())
	n := len(w.ToFind)/2 - 1
	for i := 0; i < len(w.ToFind)-1; i++ {
		w.Word = append(w.Word, "_")
	}
	for i := 0; i < n; i++ {
		nb := rand.Intn(len(w.ToFind))
		if nb == 0 {
			w.Word[nb] = string(w.ToFind[nb])
		} else if nb > 1 {
			w.Word[nb] = string(w.ToFind[nb])
		}
	}
	fmt.Print("Voici votre mots avec les lettre donner :\n")
	w.wa_to_w()
}

func DisplayArr(arr []string) {
	for _, str := range arr {
		fmt.Print(str)
	}
	fmt.Print("\n")
}

func (w *Data_Hangman) Init() {
	w.verif_level()
	w.ChoseWord()
	w.DisplayLetters()
	w.Position_init()
	w.Point =100 
	w.NbAttempts = 10
	w.Win = 0
	w.Lose = 0
}

func (w *Data_Hangman) wa_to_w() {
	fmt.Println(w.Word)
	fmt.Println(w.ToFind)
	w.Word_Display = w.Word[0]
	for i := 1; i <= len(w.Word)-1; i++ {
		fmt.Println(w.Word_Display)
		w.Word_Display += w.Word[i]
	}
}

func (w *Data_Hangman) verif_level() {
	if w.Level == "easy" {
		w.path_file = "database/Easy.txt"
	} else if w.Level == "medium" {
		w.path_file = "database/medium.txt"
	} else if w.Level == "hard" {
		w.path_file = "database/hard.txt"
	}
}



func (w *Data_Hangman) ReInit() {
	w.Word   = []string{} 
	w.ToFind = ""
	w.Word_Display =""
	w.Propo_let = []string{}
	w.NbAttempts = 10
	w.Attempts = 0
	w.ChoseWord()
	w.DisplayLetters()
}