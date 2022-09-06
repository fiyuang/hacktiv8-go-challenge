package firstChallange

import (
	"fmt"
	"strings"
)

func calcword() {
	word := "selamat malam"
	splitWord := strings.Split(word, "")

	for _, value := range splitWord {
		fmt.Println(value)
	}

	fmt.Println(WordCount(splitWord))
}

func WordCount(s []string) map[string]int {
	m := make(map[string]int)
	for _, word := range s {
		m[word] += 1
	}
	return m
}
