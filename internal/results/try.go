package results

import (
	"fmt"
)

func CountLetter(l rune, word string) int {
	var res int
	if word == "" {
		return res
	}
	for _, c := range word {
		if c == l {
			res++
		}
	}
	return res
}

// Try return the result to guess a word against another
func Try(word, guess string) (string, error) {
	var result string
	letterMap := make(map[string]int)
	if len(word) != 5 || len(guess) != 5 {
		return "", fmt.Errorf("word %s or guess %s have not the expected length (5)", word, guess)
	}
	for i, c := range guess {
		if byte(c) == word[i] {
			result += "2"
			letterMap[string(c)] = 2
		} else if n := CountLetter(c, word); n > 0 {
			_, ok := letterMap[string(c)]
			if n == 1 && ok {
				result += "0"
			} else if n == 1 && !ok {
				result += "1"
			} else if n > 1 {
				result += "1"
			}
			letterMap[string(c)] = 1
		} else {
			result += "0"
		}
	}
	return result, nil
}
