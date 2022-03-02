package main

import (
	"bufio"
	"flag"
	"os"

	"github.com/jtbonhomme/wordlebot/internal/results"
	"github.com/schollz/progressbar/v3"
	log "github.com/sirupsen/logrus"
)

func main() {
	var local = flag.String("l", "assets/words.txt", "use local words list")
	var debug = flag.Bool("d", false, "display debug information")
	flag.Parse()
	log.Infoln("start with local words list", *local)

	if *debug {
		log.SetLevel(log.DebugLevel)
	}
	file, err := os.Open(*local)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var words []string

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	bar := progressbar.Default(int64(len(words)))

	var maxEntropy float64
	var bestWord string
	for _, word := range words {
		bar.Add(1)
		g := results.New(words)
		e, err := g.Entropy(word)
		if err != nil {
			log.Panic(err)
		}
		if e > maxEntropy {
			maxEntropy = e
			bestWord = word
		}
	}
	log.Infof("Best Word to start with is %s with a entropy of %f", bestWord, maxEntropy)
}
