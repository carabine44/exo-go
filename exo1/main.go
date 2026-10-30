package main

import (
	"fmt"
	"os"
)

const MaxDisplay = 10

func main() {
	words := os.Args[1:]

	if len(words) == 0 {
		fmt.Println("Aucun mot fourni")
		os.Exit(1)
	}

	fmt.Printf("Nombre total de mots : %d\n", len(words))

	fmt.Println("Mots de longueur > 4 :")

	count := 0

	for _, word := range words {
		if len(word) > 4 {
			fmt.Println(word)
			count++

			if count >= MaxDisplay {
				break
			}
		}
	}
}
