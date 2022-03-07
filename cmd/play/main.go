package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/jtbonhomme/wordlebot/internal/wordle"
	log "github.com/sirupsen/logrus"
)

func main() {
	var upperCase = flag.Bool("c", true, "words list is in upper case (default)")
	var local = flag.String("l", "assets/long-words.txt", "use local words list")
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

	g := wordle.New(words)
	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter a word: ")
		word, err := reader.ReadString('\n')
		if err != nil {
			log.Panic(err)
		}
		word = strings.TrimSuffix(word, "\n")

		fmt.Print("Enter a result: ")
		result, err := reader.ReadString('\n')
		if err != nil {
			log.Panic(err)
		}
		result = strings.TrimSuffix(result, "\n")
		next, ent, err := g.NextWord(word, result, *upperCase)
		if err != nil {
			log.Panic(err)
		}
		log.Infof("Best word to continue with is %s with a entropy of %f", next, ent)
	}
}
