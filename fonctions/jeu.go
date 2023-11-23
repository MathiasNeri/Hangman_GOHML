package fonction

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func JouerAuPendu(nomFichier, motSecret string) {
	essaisRestants := 10
	lettresRestantes := len(motSecret) - 2
	lettresDejaEssayees := make(map[rune]bool)
	lettresDejaTrouvees := make(map[rune]bool)

	// Appeler RandomLetters pour obtenir la chaîne cachée et les indices
	motCache, indicesLettres := RandomLetters(motSecret)

	// Ajouter les deux lettres correspondant aux indices dans lettresDejaEssayees
	lettresDejaEssayees[rune(motSecret[indicesLettres[0]])] = true
	lettresDejaEssayees[rune(motSecret[indicesLettres[1]])] = true

	fmt.Println("Bienvenue au jeu du pendu!")

	var images []string
	path := "./data/image_pendu.txt"

	// Charger les images du pendu depuis le fichier
	fichierImages, err := os.Open(path)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier d'images:", err)
		return
	}
	defer fichierImages.Close()

	scannerImages := bufio.NewScanner(fichierImages)
	var image string
	for scannerImages.Scan() {
		ligne := scannerImages.Text()
		if ligne == "=========" {
			if image != "" {
				images = append([]string{image}, images...) // Ajoutez l'image en haut de la pile
			}
			image = ""
		} else {
			image += ligne + "\n"
		}
	}
	if image != "" {
		images = append([]string{image}, images...) // Ajoutez la dernière image en haut de la pile
	}

	lettreDevinee := false
	motDevine := false

	// Afficher le mot caché avant le premier tour
	fmt.Printf("Mot caché: %s\n", motCache)

	for essaisRestants > 0 && lettresRestantes > 0 && !(lettreDevinee && motDevine) {
		var entree string
		fmt.Print("Devinez une lettre ou un mot: ")
		fmt.Scan(&entree)

		if len(entree) == 1 {
			lettre := unicode.ToLower([]rune(entree)[0])
			if !unicode.IsLetter(lettre) {
				fmt.Println("Veuillez entrer une lettre alphabétique.")
				continue
			}

			if lettresDejaEssayees[lettre] {
				fmt.Println("Vous avez déjà essayé cette lettre.")
				entree = ""
				continue
			}

			lettresDejaEssayees[lettre] = true

			lettreDevinee = false // Réinitialiser la variable lettreDevinee
			if strings.ContainsRune(motSecret, lettre) {
				fmt.Println("Bonne devinette!")
				for i, char := range motSecret {
					if char == lettre {
						motCache = motCache[:i] + string(char) + motCache[i+1:]
						lettresRestantes--
					}
				}
				lettreDevinee = true // Marquer la lettre comme correctement devinée

				if lettresDejaTrouvees[lettre] {
					// Si la lettre est déjà trouvée ailleurs, révèle-la à toutes les positions
					for i, char := range motSecret {
						if char == lettre {
							motCache = motCache[:i] + string(char) + motCache[i+1:]
							lettresRestantes--
						}
					}
				}
			} else {
				fmt.Println("Mauvaise devinette!")
				essaisRestants -= 1 // Retirer 1 essai
			}
		} else if len(entree) == len(motSecret) && entree == motSecret {
			lettresRestantes = 0 // Le joueur a correctement deviné le mot
			motDevine = true
		} else {
			fmt.Println("Mauvaise devinette!")
			if len(entree) > 1 {
				essaisRestants -= 2 // Retirer 2 essais seulement si la tentative est un mot complet
			}
		}
		fmt.Printf("Mot caché: %s\n", motCache)
		fmt.Printf("Tentatives restantes: %d\n", essaisRestants)
		// Afficher les images avant la décrémentation de essaisRestants
		if essaisRestants >= 0 && essaisRestants < len(images) {
			fmt.Println(images[essaisRestants])
		} else {
			fmt.Println("Nombre d'essais non pris en charge.")
		}
	}

	if lettresRestantes == 0 || motDevine {
		fmt.Println("Félicitations, vous avez gagné! Le mot était :", motSecret)
	} else {
		fmt.Println("Désolé, vous avez perdu. Le mot était :", motSecret)
	}
}
