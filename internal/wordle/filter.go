package wordle

import (
	"strings"
)

// Filter remove words that don't match the guess.
// For example, "abcde" + 00102 will remove all words
// which contains the following letters a, b or d
// which do not contain the letter c
// and do not have a letter e in last position
func (g *Game) Filter(word string, result []int) {
	g.filteredWords = g.words
	g.discardedWords = []string{}

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

// FilteredWords return a string with the last filtered words
func (g *Game) FilteredWords() string {
	var s string
	for _, word := range g.filteredWords {
		s += word
		s += " "
	}
	return s
}

// Commit persists filtered word list as main word list
func (g *Game) Commit() {
	g.words = g.filteredWords
	g.filteredWords = []string{}
	g.discardedWords = []string{}
}

// DiscardedWords return a string with the last discarded words
func (g *Game) DiscardedWords() string {
	var s string
	for _, word := range g.discardedWords {
		s += word
		s += " "
	}
	return s
}
