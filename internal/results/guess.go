package results

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	Nothing   int = 0
	InWord    int = 1
	GoodPlace int = 2
)

// Gues is a structure that support entropy computation
type Guess struct {
	words         []string
	filteredWords []string
}

// New create a new Guess object from a given word list
func New(words []string) *Guess {
	return &Guess{
		words:         words,
		filteredWords: words,
	}
}

// Information decompose an integer into power of 3
// 0 <= i < 3^5
// i = a x 3^4 + b x 3^3 + c x 3^2 + d x 3^1 + e x 3^0
func Information(i int) []int {
	res := []int{}
	if i >= int(math.Pow(3, 5)) || i < 0 {
		return res
	}

	for n := 4; n >= 0; n-- {
		a := i / int(math.Pow(3, float64(n)))
		res = append(res, a)
		i -= a * int(math.Pow(3, float64(n)))
	}

	return res
}

func hasLetter(word string, c rune) bool {
	for i := 0; i < len(word); i++ {
		if word[i] == byte(c) {
			return true
		}
	}
	return false
}

func hasLetterInPos(word string, c rune, p int) bool {
	for i := 0; i < len(word); i++ {
		if word[i] == byte(c) && i == p {
			return true
		}
	}
	return false
}

// RemoveLetter filters all words with the letter c in any position
func (g *Guess) RemoveLetter(c rune) {
	var filteredWords []string
	for _, word := range g.filteredWords {
		if !hasLetter(word, c) {
			filteredWords = append(filteredWords, word)
		}
	}
	g.filteredWords = filteredWords
}

// RemoveNoLetter filters all words with no letter c in any position
func (g *Guess) RemoveNoLetter(c rune) {
	var filteredWords []string
	for _, word := range g.filteredWords {
		if hasLetter(word, c) {
			filteredWords = append(filteredWords, word)
		}
	}
	g.filteredWords = filteredWords
}

// RemoveLetterInPos filters all words with the letter c in position i
func (g *Guess) RemoveLetterInPos(c rune, i int) {
	var filteredWords []string
	for _, word := range g.filteredWords {
		if !hasLetterInPos(word, c, i) {
			filteredWords = append(filteredWords, word)
		}
	}
	g.filteredWords = filteredWords
}

// RemoveNoLetterInPos filters all words with no letter c in position i
func (g *Guess) RemoveNoLetterInPos(c rune, i int) {
	var filteredWords []string
	for _, word := range g.filteredWords {
		if hasLetterInPos(word, c, i) {
			filteredWords = append(filteredWords, word)
		}
	}
	g.filteredWords = filteredWords
}

// RemoveWord definitively removes a word from words list
func (g *Guess) RemoveWord(w string) {
	var filteredWords []string
	for _, word := range g.words {
		if w != word {
			filteredWords = append(filteredWords, word)
		}
	}
	g.words = filteredWords
}

// Filter remove words that don't match the guess.
// For example, "abcde" + 00102 will remove all words
// which contains the following letters a, b or d
// which do not contain the letter c
// and do not have a letter e in last position
func (g *Guess) Filter(word string, result []int) {
	g.filteredWords = g.words
	if len(word) != 5 || len(result) != 5 {
		return
	}
	word = strings.ToLower(word)
	for i, c := range word {
		switch result[i] {
		case Nothing:
			g.RemoveLetter(c)
		case InWord:
			g.RemoveNoLetter(c)
			g.RemoveLetterInPos(c, i)
		case GoodPlace:
			g.RemoveNoLetterInPos(c, i)
		}
	}
}

// Entropy compute the averge information quantity that can be retrieve
// from a word. It will check for every guess possibility and get the
// information quantity from each guess (3^5) and return the average
// information quantity.
func (g *Guess) Entropy(word string) (float64, error) {
	var entropy []float64
	var total float64
	if word == "" {
		return 0.0, fmt.Errorf("word can not be empty")
	}
	word = strings.ToLower(word)
	filename := "assets/" + word + ".stat"
	stat, err := os.Create(filename)
	if err != nil {
		return 0.0, fmt.Errorf("can not open %s: %w", filename, err)
	}
	defer stat.Close()

	for i := 0; i < int(math.Pow(3, 5)); i++ {
		result := Information(i)
		g.Filter(word, result)
		var iqty float64
		if len(g.filteredWords) != 0 {
			iqty = -math.Log(float64(len(g.filteredWords))/float64(len(g.words))) / math.Log(3)
		}
		entropy = append(entropy, iqty)
		_, err := stat.Write([]byte(fmt.Sprintf("%d%d%d%d%d,%f\n", result[0], result[1], result[2], result[3], result[4], iqty)))
		if err != nil {
			return 0.0, fmt.Errorf("can not write into %s: %w", filename, err)
		}
	}
	for _, e := range entropy {
		total += e
	}
	meanEntropy := total / float64(len(entropy))
	return meanEntropy, nil
}

// ToString return a string with all dictionary
func (g *Guess) ToString() string {
	var s string
	for _, word := range g.words {
		s += word
		s += " "
	}
	return s
}

// FilteredToString return a string with the current filtered dictionary
func (g *Guess) FilteredToString() string {
	var s string
	for _, word := range g.filteredWords {
		s += word
		s += " "
	}
	return s
}

// NextWord finds, given a game state (last proposed word and result)
// the next best word to be played
func (g *Guess) NextWord(word, res string) (string, float64, error) {
	var result []int
	if len(word) != 5 {
		return "", 0.0, fmt.Errorf("word %s has not the right length", word)
	}

	word = strings.ToLower(word)

	for _, c := range res {
		i, err := strconv.Atoi(string(c))
		if err != nil {
			return "", 0.0, fmt.Errorf("failed to convert %s into int", string(c))
		}
		result = append(result, i)
	}

	if len(result) != 5 {
		return "", 0.0, fmt.Errorf("result %v has not the right length", result)
	}

	g.Filter(word, result)
	var maxEntropy float64
	var bestWord string
	for _, w := range g.filteredWords {
		e, err := g.Entropy(w)
		if err != nil {
			log.Panic(err)
		}
		if e > maxEntropy {
			maxEntropy = e
			bestWord = w
		}
	}
	return bestWord, maxEntropy, nil
}

// Try return the result to guess a word against another
func (g *Guess) Try(word, guess string) (string, error) {
	var result string
	if len(word) != 5 || len(guess) != 5 {
		return "", fmt.Errorf("word %s or guess %s have not the right length", word, guess)
	}
	for i, c := range guess {
		if byte(c) == word[i] {
			result += "2"
		} else if strings.Contains(word, string(c)) {
			result += "1"
		} else {
			result += "0"
		}
	}
	return result, nil
}
