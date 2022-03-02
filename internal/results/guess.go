package results

import (
	"fmt"
	"math"
	"os"
	"strings"
)

const (
	Nothing   int = 0
	InWord    int = 1
	GoodPlace int = 2
)

// Gues is a structure that support entropy computation
type Guess struct {
	words          []string
	filteredWords  []string
	discardedWords []string
}

// New create a new Guess object from a given word list
func New(words []string) *Guess {
	return &Guess{
		words:         words,
		filteredWords: words,
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
