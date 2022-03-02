package results

import (
	"fmt"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

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
	log.Debugf("\t%d filtered words are: %s", len(g.filteredWords), g.FilteredToString())
	log.Debugf("\t%d discarded words are: %s", len(g.discardedWords), g.DiscardedToString())
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
