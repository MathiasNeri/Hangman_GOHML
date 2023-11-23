package main

import (
	"html/template"
	"net/http"
)

func main() {
	// Gestion de tous les fichiers gohtml
	tmpl := template.Must(template.ParseGlob("templates/*.gohtml"))

	// Gestion de tous les fichiers css
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Gestion de tous les fichiers png (Hangman + icone de page)
	images := http.FileServer(http.Dir("images"))
	http.Handle("/images/", http.StripPrefix("/images/", images))

	// Initialisation d'une variable nécessaire au reset du mot
	check := 0

	// Gestion de la page d'accueil
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		check = 0
		easy = Hangman{Lives: 10, Win: false, Loose: false, File: "./text/easy.txt"}
		normal = Hangman{Lives: 10, Win: false, Loose: false, File: "./text/normal.txt"}
		hard = Hangman{Lives: 10, Win: false, Loose: false, File: "./text/hard.txt"}
		err := tmpl.ExecuteTemplate(w, "index", "")
		if err != nil {
			return
		}
	})

	// Gestion de la page de jeu en mode facile
	http.HandleFunc("/hangmanEasy", func(w http.ResponseWriter, r *http.Request) {
		for check == 0 {
			easy.getRandomWord()
			check += 1
		}
		easy.start(r)
		err := tmpl.ExecuteTemplate(w, "hangmanEasy", easy)
		if err != nil {
			return
		}
	})

	// Gestion de la page de jeu en mode normal
	http.HandleFunc("/hangman", func(w http.ResponseWriter, r *http.Request) {
		for check == 0 {
			normal.getRandomWord()
			check += 1
		}
		normal.start(r)
		err := tmpl.ExecuteTemplate(w, "hangman", normal)
		if err != nil {
			return
		}
	})

	// Gestion de la page de jeu en mode difficile
	http.HandleFunc("/hangmanHard", func(w http.ResponseWriter, r *http.Request) {
		for check == 0 {
			hard.getRandomWord()
			check += 1
		}
		hard.start(r)
		err := tmpl.ExecuteTemplate(w, "hangmanHard", hard)
		if err != nil {
			return
		}
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
