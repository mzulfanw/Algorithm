package main

import (
	"fmt"
	"strings"
)

func main() {

	text := "Belajar golang dengan mudah"

	vocab := [5]string{"a", "i", "u", "e", "o"}

	var tempWord string
	for i := 0; i < len(text); i++ {
		tempWord = strings.ToLower(string(text[i]))
		for j := 0; j < len(vocab); j++ {
			if tempWord == vocab[j] {
				fmt.Println("index", i, "word", tempWord)
			}
		}
	}
}
