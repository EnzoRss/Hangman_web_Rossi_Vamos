package database

import (
	"fmt"
	"os"
	"strings"
)

func (w *Data_Hangman) Input(input string) {
	if VerifInput(input) && !VerifArr(w.Propo_let, input) {
		if len(input) >= 2 {
			if w.VerifWord(input) {
				fmt.Println("vous avez gagner !!!")
				w.Result = "Win"
				w.Win++
			}
		} else if len(input) == 1 {
			w.VerifLetter(input)
			if !w.VerifVictory() {
				fmt.Println("vous avez gagner !!!")
				w.Result = "Win"
				w.Win++
			}
		}
		fmt.Println(w.HangmanPositions[w.Attempts])

		fmt.Println("il vous reste encore ", 10-w.Attempts, " d'essaie")
		w.NbAttempts = 10 - w.Attempts
	} else if input == "STOP" {
		w.saveData()
	} else if VerifArr(w.Propo_let, input) {
		fmt.Println("Vous avez déja proposer ce mot")
		w.Result = "Vous avez déja proposer ce mot"
	} else {
		fmt.Println("vous n'avez pas rentrée un charactère acceptable ")
	}
	if w.Attempts == 10 {
		fmt.Println("Vous avez PERDU !!!")
		w.Result = "lose"
		w.Lose++
	}

	if !VerifArr(w.Propo_let, input) {
		w.Propo_let = append(w.Propo_let, input)
	}
	w.wa_to_w()
}

func (w *Data_Hangman) VerifLetter(str string) {
	var temp bool
	verify := false
	var compt int
	for index, letter := range w.ToFind {
		if str == string(letter) {
			for _, verif := range w.Word {
				if string(letter) == verif {
					temp = true
				}
				if temp && compt <= w.VerifNbLetter(str) {
					fmt.Println("La lettre est déja présente")
					w.Result = "La lettre est déja présente"
					break
				} else {
					w.Word[index] = string(letter)
					compt++
					w.Point += 10
					verify = true
				}
			}
		}
	}
	if verify && !temp {
		fmt.Println("Vous avez ajouter une lettre")
		w.Result = "Vous avez ajouter une lettre"
	} else if !verify && !temp {
		fmt.Println("La lettre rentrée n'est pas dans le mots ")
		w.Result = "La lettre rentrée n'est pas dans le mot"
		w.Attempts++
		w.Point -= 10
	}
	fmt.Println(w.Word)
}

func (w *Data_Hangman) VerifWord(str string) bool {
	temp := true
	for i := range str {
		if str[i] != w.ToFind[i+1] {
			temp = false
		}
	}
	if temp {
		fmt.Println("vous avez trouver le mot")
		w.Result = "vous avez trouver le mot"
		return true
	} else {
		fmt.Println("Vous vous êtes trompé ")
		w.Attempts = w.Attempts + 2
		fmt.Println(w.Attempts)
		fmt.Println(w.Word)
		return false
	}
}

func (w Data_Hangman) VerifNbLetter(str string) int {
	var compt int
	for _, letter := range str {
		if string(letter) == w.ToFind {
			compt++
		}
	}
	return compt
}

func (w *Data_Hangman) Position_init() {
	arr, _ := os.ReadFile("database/hangman.txt")
	arrS := strings.Split(string(arr), "=========")
	for i := range arrS {
		w.HangmanPositions[i] = arrS[i]
	}
	for i := 0; i < len(w.HangmanPositions)-1; i++ {
		w.HangmanPositions[i] += "========="
	}

}

func (w Data_Hangman) VerifVictory() bool {
	var temp bool
	for i := 0; i < len(w.Word); i++ {
		if w.Word[i] == "_" {
			temp = true
		}
	}

	return temp
}

func VerifInput(str string) bool {
	var temp bool
	for _, letter := range str {
		if letter >= 7 && letter <= 122 {
			temp = true

		} else {
			temp = false

		}
	}
	return temp
}

func VerifArr(arr []string, str string) bool {
	verif := false
	for _, letter := range arr {
		if letter == str {
			verif = true
		}
	}
	return verif
}