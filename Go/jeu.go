package main

import (
	"net/http"
)

// Hangman Structure nécessaire au déroulement du jeu
type Hangman struct {
	WordToGuess  string
	HiddenWord   []string
	UserInput    string
	Lives        int
	Proposition  []string
	FoundLetters int
	Win          bool
	Loose        bool
	File         string
}

// Initialisation des différentes variables de jeu pour chaque difficulté
var easy = Hangman{}
var hard = Hangman{}
var normal = Hangman{}

// Lancement du jeu Hangman
func (user *Hangman) start(r *http.Request) { //Programme de lancement du jeu
	verifyLettersUsed := 0
	verifyGoodProposition := 0
	user.UserInput = r.FormValue("userinput")
	verifyLettersUsed = alreadyUsed(user, verifyLettersUsed)
	addProp(verifyLettersUsed, user)
	verifyGoodProposition = isPropTrue(user, verifyGoodProposition)
	win(user)
	livesChange(verifyGoodProposition, user)
	loose(user)
}

// Vérification si la proposition effectuée a déjà été faite ou non
func alreadyUsed(user *Hangman, verifyLettersUsed int) int {
	for i := range user.Proposition {
		if user.UserInput == user.Proposition[i] {
			verifyLettersUsed++
		}
	}
	return verifyLettersUsed
}

// Ajout de la proposition à la liste des propositions passées
func addProp(verifyLettersUsed int, user *Hangman) {
	if verifyLettersUsed == 0 {
		user.Proposition = append(user.Proposition, user.UserInput)
	}
}

// Vérification si la lettre proposée est présente dans le mot
func isPropTrue(user *Hangman, verifyGoodProposition int) int {
	for i := 0; i < len(user.WordToGuess); i++ {
		if user.UserInput == string(user.WordToGuess[i]) && string(user.HiddenWord[i]) == "_" {
			user.HiddenWord[i] = string(user.WordToGuess[i])
			user.FoundLetters++
		} else {
			verifyGoodProposition++
		}
	}
	return verifyGoodProposition
}

// Validation de la victoire du joueur
func win(user *Hangman) {
	if user.UserInput == user.WordToGuess {
		user.Win = true
	}
	if user.FoundLetters == len(user.WordToGuess) {
		user.Win = true
	}
}

// Modification du compteur d'essai si la proposition est fausse
func livesChange(verifyGoodProposition int, user *Hangman) {
	if verifyGoodProposition == len(user.WordToGuess) {
		if len(user.UserInput) == 1 {
			user.Lives--
		} else if len(user.UserInput) > 1 {
			user.Lives -= 2
			if user.Lives < 0 {
				user.Lives = 0
			}
		} else {
		}
	}
}

// Validation de la défaite du joueur
func loose(user *Hangman) {
	if user.Lives <= 0 {
		user.Loose = true
	}
}
