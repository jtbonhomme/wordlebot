package wordle

import (
	"fmt"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

// NextWord finds, given a game state (last proposed word and result)
// the next best word to be played
func (g *Game) NextWord(word, res string, upperCase bool) (string, float64, error) {
	var result []int
	if len(word) != 5 {
		return "", 0.0, fmt.Errorf("word %s has not the right length", word)
	}

	if upperCase {
		word = strings.ToUpper(word)
	} else {
		word = strings.ToLower(word)
	}

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

	g.Filter(word, result, upperCase)
	log.Debugf("\t%d filtered words are: %s", len(g.filteredWords), g.FilteredWords())
	g.Commit()

	var maxEntropy float64
	var bestWord string
	if len(g.words) == 0 {
		return "", 0.0, nil
	}
	for _, w := range g.words {
		e, _, err := g.Entropy(w, upperCase)
		if err != nil {
			log.Panic(err)
		}
		if e > maxEntropy {
			maxEntropy = e
			bestWord = w
		}
	}
	if len(g.words) > 0 && bestWord == "" && maxEntropy == 0 {
		return g.words[0], 0.0, nil
	}

	return bestWord, maxEntropy, nil
}
