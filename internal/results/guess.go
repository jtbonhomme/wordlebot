package results

import (
	"fmt"
	"math"
	"sort"
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

type Stat struct {
	Result  string
	Entropy float64
}

// ByEntropy implements sort.Interface for []Stat based on
// the Entropy field.
type ByEntropy []Stat

func (a ByEntropy) Len() int           { return len(a) }
func (a ByEntropy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByEntropy) Less(i, j int) bool { return a[i].Entropy < a[j].Entropy }

// Entropy compute the averge information quantity that can be retrieve
// from a word. It will check for every guess possibility and get the
// information quantity from each guess (3^5) and return the average
// information quantity.
func (g *Guess) Entropy(word string) (float64, []Stat, error) {
	var stats []Stat

	var entropy []float64
	var total float64
	if word == "" {
		return 0.0, stats, fmt.Errorf("word can not be empty")
	}
	word = strings.ToLower(word)

	for i := 0; i < int(math.Pow(3, 5)); i++ {
		result := Information(i)
		g.Filter(word, result)
		var iqty float64
		if len(g.filteredWords) != 0 {
			iqty = -math.Log(float64(len(g.filteredWords))/float64(len(g.words))) / math.Log(3)
		}
		entropy = append(entropy, iqty)
		value := Stat{
			Result:  fmt.Sprintf("%d%d%d%d%d", result[0], result[1], result[2], result[3], result[4]),
			Entropy: iqty,
		}
		stats = append(stats, value)
	}
	for _, e := range entropy {
		total += e
	}
	meanEntropy := total / float64(len(entropy))

	// Sort stats
	sort.Sort(ByEntropy(stats))
	return meanEntropy, stats, nil
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
