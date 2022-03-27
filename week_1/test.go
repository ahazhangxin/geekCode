package main

import "fmt"

func longestWord(words []string) string {
	maxlenght, res := 0, ""
	for _, word := range words {
		if maxlenght < len(word) {
			maxlenght = len(word)
			res = word
		}
	}
	return res
}

func main() {
	a := []string{"a", "banana", "app", "appl", "ap", "apply", "apple"}
	fmt.Println(longestWord(a))
}
