package fonction

import (
	"crypto/rand"
	"math/big"
	"strings"
)

func RandomLetters(maChaine string) (string, [2]int) {
	// Créer une nouvelle chaîne de caractères avec des tirets pour cacher les lettres
	chaineCachee := ""
	for j := 0; j < len(maChaine); j++ {
		chaineCachee += "-"
	}

	// Générer deux indices aléatoires pour extraire deux lettres aléatoires de la chaîne
	indices := make([]int, 0)

	for len(indices) < 2 {
		indice, _ := rand.Int(rand.Reader, big.NewInt(int64(len(maChaine))))
		i := int(indice.Int64())
		// Vérifier si l'indice n'a pas déjà été choisi et si la lettre correspondante ne se répète pas dans le mot
		if !contains(indices, i) && strings.Count(maChaine, string(maChaine[i])) < 2 {
			indices = append(indices, i)
		}
	}

	// Remplacer les tirets par les lettres aléatoires aux indices correspondants
	for _, i := range indices {
		chaineCachee = chaineCachee[:i] + string(maChaine[i]) + chaineCachee[i+1:]
	}

	// Stocker les indices dans un tableau
	indexLettres := [2]int{indices[0], indices[1]}

	// Retourner la chaîne cachée avec les lettres aléatoires et les indices
	return chaineCachee, indexLettres
}

func contains(arr []int, elem int) bool {
	for _, e := range arr {
		if e == elem {
			return true
		}
	}
	return false
}
