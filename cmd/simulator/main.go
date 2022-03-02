package main

import (
	"bufio"
	"flag"
	"os"

	"github.com/fatih/color"
	"github.com/jtbonhomme/wordlebot/internal/results"
	log "github.com/sirupsen/logrus"
)

const (
	maxAttempts int    = 6
	firstWord   string = "taris"
)

func main() {
	var local = flag.String("l", "assets/words.txt", "use local words list")
	var debug = flag.Bool("d", false, "display debug information")
	flag.Parse()
	log.Infoln("start with local words list", *local)
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

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
	var successes []int

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	// start a new game for each word and count attempts
	for _, word := range words {
		var win bool
		var attempts int
		var result string
		var err error

		log.Debugf("Try to guess word %s", word)
		lastWord := "taris"
		g := results.New(words)

		// first attempt with "taris"
		result, err = g.Try(word, lastWord)
		if err != nil {
			log.Panic(err)
		}
		log.Debugf("\t[%d] guess: %s result: %s", attempts, green(lastWord), red(result))
		attempts++
		// Did we win with the first attempts?
		if result == "22222" {
			win = true
		}
		// next attempts
		for ; attempts < maxAttempts && !win; attempts++ {
			nextWord, _, err := g.NextWord(lastWord, result)
			if err != nil {
				log.Panic(err)
			}
			result, err = g.Try(word, nextWord)
			if err != nil {
				log.Panic(err)
			}
			log.Debugf("\t[%d] guess: %s result: %s", attempts, green(nextWord), red(result))
			if result == "22222" {
				win = true
			}
			lastWord = nextWord
		}
		if win {
			log.Debugf("%s ✅ Found word %s in %d attempts", green("SUCCESS"), word, attempts)
			successes = append(successes, attempts)
		} else {
			log.Debugf("%s ❌ Couldn't find word %s in %d attempts or less", red("FAILURE"), word, maxAttempts)
		}
	}
	var averageAttempts float64
	for _, a := range successes {
		averageAttempts += float64(a)
	}
	averageAttempts /= float64(len(successes))
	log.Infof("wordlebot performance is %f attempts to guess a word", averageAttempts)
}
