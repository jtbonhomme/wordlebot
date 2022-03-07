package wordle

import (
	"fmt"
	"strings"
)

// CountLetter return the occurrence of a letter in a word
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

// Try return the comparison of a word against another
// Each letter of the first word is compared against the
// letter at the same place in the second word.
// For each letter, a code is provided:
// - 0 : this letter do not correspond
// - 1 : this letter exist in the second word in an other place
// - 2 : this letter is the same in both word at the same place
// Ex.:
// Try("abcde", "afdge") will return "20102"
func Try(word, guess string, upperCase bool) (string, error) {
	var result string
	letterMap := make(map[string]int)
	if len(word) != 5 || len(guess) != 5 {
		return "", fmt.Errorf("word %s or guess %s have not the expected length (5)", word, guess)
	}

	if upperCase {
		word = strings.ToUpper(word)
		guess = strings.ToUpper(guess)
	} else {
		word = strings.ToLower(word)
		guess = strings.ToLower(guess)
	}

	for i, c := range guess {
		if byte(c) == word[i] {
			result += GoodPlaceStr
			letterMap[string(c)] = 2
		} else if n := CountLetter(c, word); n > 0 {
			_, ok := letterMap[string(c)]
			if n == 1 && ok {
				result += NothingStr
			} else if n == 1 && !ok {
				result += InWordStr
			} else if n > 1 {
				result += InWordStr
			}
			letterMap[string(c)] = 1
		} else {
			result += NothingStr
		}
	}
	return result, nil
}
