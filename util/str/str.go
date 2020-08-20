package str

import (
	"strings"
)

func SnakeToCamel(word string) string {
	words := strings.Split(strings.Trim(word, "_"), "_")
	for i, w := range words {
		if i != 0 {
			words[i] = strings.ToUpper(w[:1]) + w[1:]
		}
	}
	return strings.Join(words, "")
}

func SnakeToBigCamel(word string) string {
	words := strings.Split(strings.Trim(word, "_"), "_")
	for i, w := range words {
		words[i] = strings.ToUpper(w[:1]) + w[1:]
	}
	return strings.Join(words, "")
}

func CamelToSnake(word string) string {
	var words []byte
	for k, w := range word {
		if 'A' <= w && w <= 'Z' {
			if 0 < k {
				words = append(words, '_')
			}
			words = append(words, byte(w)+32)
		} else {
			words = append(words, byte(w))
		}
	}
	return string(words)
}
