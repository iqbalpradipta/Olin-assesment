package main

import (
	"fmt"
	"strings"
)

// Belum selesai

func substring(words []string, s string) []int {
	var stringArr []string
	wordsJoin := strings.Join(words, "")
	wordsSplit := strings.Split(wordsJoin, "")
	sSplit := strings.Split(s, "")
	for i := 0; i < len(sSplit); i++ {
		if sSplit[i] == wordsSplit[i] {
			stringArr = append(stringArr, sSplit[i])
		}
	}
	fmt.Println(stringArr)
	return []int{}
}

func main() {
	words := []string{"word","good","best","word"}
	s := "wordgoodgoodgoodbestword"

	fmt.Print(substring(words, s))
}


