package main

import (
	"bufio"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var (
	scoreboard int
	nberror    int
	difficulty string
	words      []string
	guesses    []string
)

func main() {
	temp, err := template.ParseGlob("./templates/*.html")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("test")
		temp.ExecuteTemplate(w, "index", nil)
	})

	http.HandleFunc("/choix", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "choix", nil)
	})

	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/asset"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	http.HandleFunc("/sauvegarde", func(w http.ResponseWriter, r *http.Request) {
		difficulty = r.FormValue("difficulty")
	})

	http.HandleFunc("/game", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "game", difficulty)
	})

	http.HandleFunc("/result", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "result", r)
	})

	http.HandleFunc("/getword", func(w http.ResponseWriter, r *http.Request) {
		word, err := GetWord(difficulty + ".txt")
		if err != nil {
			http.Error(w, "Failed to get word", http.StatusInternalServerError)
			return
		}
		w.Write([]byte(word))
	})

	// Your existing server setup...
	http.ListenAndServe(":8080", nil)
}

func GetWord(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if len(words) == 0 {
		return "", fmt.Errorf("no words available")
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIndex := r.Intn(len(words))

	return words[randomIndex], nil
}

func Display() {
	// Your display logic...
}

func Verifier(guess string) {
	if len(guess) == 1 { // The player guessed a letter
		// Check if the guessed letter is correct
		guessed := false
		for _, letter := range words {
			if letter == guess {
				guessed = true
				break
			}
		}
		if !guessed {
			nberror++
		} else {
			guesses = append(guesses, guess)
		}
	}
}
