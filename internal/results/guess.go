package results

import (
	"math"

	log "github.com/sirupsen/logrus"
)

const (
	Nothing   int = 0
	InWord    int = 1
	GoodPlace int = 2
)

type Guess struct {
	words []string
}

func New(words []string) *Guess {
	return &Guess{
		words: words,
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
	for _, word := range g.words {
		if !hasLetter(word, c) {
			log.Debugf("keep word %s because it does not have letter %s", word, string(c))
			filteredWords = append(filteredWords, word)
		}
	}
	g.words = filteredWords
}

func (g *Guess) RemoveNoLetter(c rune) {
	var filteredWords []string
	for _, word := range g.words {
		if hasLetter(word, c) {
			log.Debugf("keep word %s because it has letter %s", word, string(c))
			filteredWords = append(filteredWords, word)
		}
	}
	g.words = filteredWords
}

func (g *Guess) RemoveLetterInPos(c rune, i int) {
	var filteredWords []string
	for _, word := range g.words {
		if !hasLetterInPos(word, c, i) {
			log.Debugf("keep word %s because it does not have letter %s in position %d", word, string(c), i)
			filteredWords = append(filteredWords, word)
		}
	}
	g.words = filteredWords
}

func (g *Guess) RemoveNoLetterInPos(c rune, i int) {
	var filteredWords []string
	for _, word := range g.words {
		if hasLetterInPos(word, c, i) {
			log.Debugf("keep word %s because it has letter %s", word, string(c))
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
	if len(word) != 5 || len(result) != 5 {
		return
	}
	for i, c := range word {
		switch result[i] {
		case Nothing:
			log.Debugf("remove words that contain the letter %s", string(c))
			g.RemoveLetter(c)
		case InWord:
			log.Debugf("remove words that do not contain the letter %s and words that contains the letter %s at the %dth position", string(c), string(c), i+1)
			g.RemoveNoLetter(c)
			g.RemoveLetterInPos(c, i)
		case GoodPlace:
			log.Debugf("remove words that do not contain the letter %s at the %dth position", string(c), i+1)
			g.RemoveNoLetterInPos(c, i)
		}
	}
}

func (g *Guess) Entropy(s string, result []int) float64 {
	return 1.0
}
