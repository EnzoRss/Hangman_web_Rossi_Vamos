package database

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (w *Data_Hangman) Input() {
	for w.Attempts < 10 {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("Ecrit la lettre ou le mots de ton choix : ")
		scanner.Scan()
		input := scanner.Text()
		if VerifInput(input) {
			if len(input) >= 2 {
				w.VerifWord(input)
				if w.VerifWord(input) {
					fmt.Println("vous avez gagner !!!")
					break
				}
			} else if len(input) == 1 {
				w.VerifLetter(input)
				if !w.VerifVictory() {
					fmt.Println("vous avez gagner !!!")
					break
				}
			}
			fmt.Println(w.HangmanPositions[w.Attempts])

			fmt.Println("il vous reste encore ", 10-w.Attempts, " d'essaie")

		} else if input == "STOP" {
			w.saveData()
			break
		} else {
			fmt.Println("vous n'avez pas rentrée un chractère acceptable ")
		}

	}
	if w.Attempts == 10 {
		fmt.Println("Vous avez PERDU !!!")
	}
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
					break
				} else {
					w.Word[index-1] = string(letter)
					compt++
					verify = true
				}
			}
		}
	}
	if verify && !temp {
		fmt.Println("Vous avez ajouter une lettre")
	} else if !verify && !temp {
		fmt.Println("La lettre rentrée n'est pas dans le mots ")
		w.Attempts++
	}
	fmt.Println(w.Word)
}

func (w *Data_Hangman) VerifWord(str string) bool {
	temp := true
	for i := 0; i < len(w.ToFind)-1; i++ {
		if str[i] != w.ToFind[i+1] {
			temp = false
		}
	}
	if temp {
		fmt.Println("vous avez trouver le mot")
		return true
	} else {
		fmt.Println("Vous vous êtes trompé ")
		w.Attempts += 2
		fmt.Println(w.Word)
		return false
	}
}

func (w Data_Hangman) VerifNbLetter(str string) int {
	var compt int
	for _, letter := range w.ToFind {
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
		if letter >= 97 && letter <= 122 {
			temp = true
		} else {
			temp = false
		}
	}
	return temp
}
