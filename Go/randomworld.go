package main

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"
)

// Choix du mot aléatoirement dans le dossier .txt
func (user *Hangman) getRandomWord() {
	var array []string
	fileScanner := createScanner(user.File)
	array = getWords(fileScanner, array)
	rand.Seed(time.Now().UnixNano())
	ran := rand.Intn(len(array))
	user.WordToGuess = array[ran]
	user.HiddenWord = hideToFindWord(user.WordToGuess)
	user.FoundLetters = user.showToFindLetters()
}

// Création du scanner pour lire les fichiers txt
func createScanner(fileName string) *bufio.Scanner { //Programme de création d'un scanner
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	fileScanner := bufio.NewScanner(file)
	return fileScanner
}

// Récupération des mots des fichiers txt
func getWords(fileScanner *bufio.Scanner, array []string) []string {
	for fileScanner.Scan() {
		array = append(array, fileScanner.Text())
	}
	return array
}

// Création du mot caché ("______")
func hideToFindWord(word string) []string {
	var hiddenWord []string
	for i := 0; i < len(word); i++ {
		hiddenWord = append(hiddenWord, "_")
	}
	return hiddenWord
}

// Choix des lettres qui sont affichées dès le début
func (user *Hangman) showToFindLetters() int {
	lettersToDisplay := (len(user.HiddenWord) / 2) - 1
	var displayedLetters int
	for i := 0; i < lettersToDisplay; i++ {
		index := rand.Intn(len(user.HiddenWord))
		if user.HiddenWord[index] == "_" {
			displayedLetters++
		}
		user.HiddenWord[index] = string(user.WordToGuess[index])
	}
	return displayedLetters
}
