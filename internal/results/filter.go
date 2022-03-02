package results

import (
	"strings"
)

// Filter remove words that don't match the guess.
// For example, "abcde" + 00102 will remove all words
// which contains the following letters a, b or d
// which do not contain the letter c
// and do not have a letter e in last position
func (g *Guess) Filter(word string, result []int) {
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

// FilteredToString return a string with the last filtered words
func (g *Guess) FilteredToString() string {
	var s string
	for _, word := range g.filteredWords {
		s += word
		s += " "
	}
	return s
}

// Commit persists filtered word list as main word list
func (g *Guess) Commit() {
	g.words = g.filteredWords
	g.filteredWords = []string{}
	g.discardedWords = []string{}
}

// DiscardedToString return a string with the last discarded words
func (g *Guess) DiscardedToString() string {
	var s string
	for _, word := range g.discardedWords {
		s += word
		s += " "
	}
	return s
}
