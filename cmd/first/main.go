package main

import (
	"bufio"
	"flag"
	"os"

	"github.com/jtbonhomme/wordlebot/internal/results"
	log "github.com/sirupsen/logrus"
)

func main() {
	var local = flag.String("l", "assets/words.txt", "use local words list")
	var debug = flag.Bool("d", false, "display debug information")
	flag.Parse()
	log.Infoln("start with local words list ", *local)

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

	g := results.New(words)

	e1, _ := g.Entropy("tarie")
	e2, _ := g.Entropy("tarin")
	e3, _ := g.Entropy("round")
	log.Debugf("tarie entropy: %f", e1)
	log.Debugf("tarin entropy: %f", e2)
	log.Debugf("round entropy: %f", e3)

}
