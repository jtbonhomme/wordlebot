package wordle

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
func (g *Game) RemoveLetter(c rune) {
	var filteredWords []string
	var discardedWords []string
	for _, word := range g.filteredWords {
		if !hasLetter(word, c) {
			filteredWords = append(filteredWords, word)
		} else {
			discardedWords = append(discardedWords, word)
		}
	}
	g.filteredWords = filteredWords
	g.discardedWords = append(g.discardedWords, discardedWords...)
}

// RemoveNoLetter filters all words with no letter c in any position
func (g *Game) RemoveNoLetter(c rune) {
	var filteredWords []string
	var discardedWords []string
	for _, word := range g.filteredWords {
		if hasLetter(word, c) {
			filteredWords = append(filteredWords, word)
		} else {
			discardedWords = append(discardedWords, word)
		}
	}
	g.filteredWords = filteredWords
	g.discardedWords = append(g.discardedWords, discardedWords...)
}

// RemoveLetterInPos filters all words with the letter c in position i
func (g *Game) RemoveLetterInPos(c rune, i int) {
	var filteredWords []string
	var discardedWords []string
	for _, word := range g.filteredWords {
		if !hasLetterInPos(word, c, i) {
			filteredWords = append(filteredWords, word)
		} else {
			discardedWords = append(discardedWords, word)
		}
	}
	g.filteredWords = filteredWords
	g.discardedWords = append(g.discardedWords, discardedWords...)
}

// RemoveNoLetterInPos filters all words with no letter c in position i
func (g *Game) RemoveNoLetterInPos(c rune, i int) {
	var filteredWords []string
	var discardedWords []string
	for _, word := range g.filteredWords {
		if hasLetterInPos(word, c, i) {
			filteredWords = append(filteredWords, word)
		} else {
			discardedWords = append(discardedWords, word)
		}
	}
	g.filteredWords = filteredWords
	g.discardedWords = append(g.discardedWords, discardedWords...)
}

// RemoveWord definitively removes a word from the main list
func (g *Game) RemoveWord(w string) {
	var filteredWords []string
	for _, word := range g.words {
		if w != word {
			filteredWords = append(filteredWords, word)
		}
	}
	g.words = filteredWords
	g.filteredWords = filteredWords
}
