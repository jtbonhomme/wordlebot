package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/jtbonhomme/wordlebot/internal/wordle"
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
		g := wordle.New(words)
		e, stats, err := g.Entropy(word)
		if err != nil {
			log.Panic(err)
		}
		if e > maxEntropy {
			maxEntropy = e
			bestWord = word
		}
		filename := "assets/" + word + ".stat"
		statFile, err := os.Create(filename)
		if err != nil {
			log.Panic(err)
		}

		for _, stat := range stats {
			_, err := statFile.Write([]byte(fmt.Sprintf("%s,%f\n", stat.Result, stat.Entropy)))
			if err != nil {
				log.Panic(err)
			}
		}

		statFile.Close()
	}

	log.Infof("Best Word to start with is %s with a entropy of %f", bestWord, maxEntropy)
}
