package results

import (
	"fmt"
	"math"
	"os"

	log "github.com/sirupsen/logrus"
)

const (
	Nothing   int = 0
	InWord    int = 1
	GoodPlace int = 2
)

type Guess struct {
	words         []string
	filteredWords []string
}

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

func (g *Guess) RemoveLetter(c rune) {
	var filteredWords []string
	for _, word := range g.filteredWords {
		if !hasLetter(word, c) {
			filteredWords = append(filteredWords, word)
		}
	}
	g.filteredWords = filteredWords
}

func (g *Guess) RemoveNoLetter(c rune) {
	var filteredWords []string
	for _, word := range g.filteredWords {
		if hasLetter(word, c) {
			filteredWords = append(filteredWords, word)
		}
	}
	g.filteredWords = filteredWords
}

func (g *Guess) RemoveLetterInPos(c rune, i int) {
	var filteredWords []string
	for _, word := range g.filteredWords {
		if !hasLetterInPos(word, c, i) {
			filteredWords = append(filteredWords, word)
		}
	}
	g.filteredWords = filteredWords
}

func (g *Guess) RemoveNoLetterInPos(c rune, i int) {
	var filteredWords []string
	for _, word := range g.filteredWords {
		if hasLetterInPos(word, c, i) {
			filteredWords = append(filteredWords, word)
		}
	}
	g.filteredWords = filteredWords
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

func (g *Guess) Entropy(word string) (float64, error) {
	var entropy []float64
	var total float64
	if word == "" {
		return 0.0, fmt.Errorf("word can not be empty")
	}

	filename := "assets/" + word + ".stat"
	stat, err := os.Create(filename)
	if err != nil {
		return 0.0, fmt.Errorf("can not open %s: %w", filename, err)
	}
	defer stat.Close()

	for i := 0; i < int(math.Pow(3, 5)); i++ {
		result := Information(i)
		g.Filter(word, result)
		e := float64(len(g.filteredWords)) / float64(len(g.words))
		entropy = append(entropy, e)
		_, err := stat.Write([]byte(fmt.Sprintf("%d%d%d%d%d,%f\n", result[0], result[1], result[2], result[3], result[4], e)))
		if err != nil {
			return 0.0, fmt.Errorf("can not write into %s: %w", filename, err)
		}
	}
	for _, e := range entropy {
		total += e
	}
	meanEntropy := total / float64(len(entropy))
	log.Debugf("entropy of %s is %f", word, meanEntropy)
	return meanEntropy, nil
}

func (g *Guess) ToString() string {
	var s string
	for _, word := range g.words {
		s += word
		s += " "
	}
	return s
}

func (g *Guess) FilteredToString() string {
	var s string
	for _, word := range g.filteredWords {
		s += word
		s += " "
	}
	return s
}
