package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/fatih/color"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/jtbonhomme/wordlebot/internal/wordle"
	log "github.com/sirupsen/logrus"
)

const (
	maxAttempts      int    = 6
	defaultFirstWord string = "SORTE"
	resultOK         string = "22222"
)

func main() {
	var progress int
	var results []int
	var firstWord = flag.String("f", defaultFirstWord, "first word (default is )")
	var upperCase = flag.Bool("c", true, "words list is in upper case (default)")
	var max = flag.String("m", "", "max words to test")
	var local = flag.String("l", "assets/words.txt", "use local words list")
	var debug = flag.Bool("d", false, "display debug information")
	flag.Parse()
	log.Infoln("start with local words list", *local)
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	// init results
	for i := 0; i < 7; i++ {
		results = append(results, 0)
	}

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
	wordsToGame := words
	maxWords := len(words)
	log.Infof("there are %d possible words", len(words))

	if *max != "" {
		m, err := strconv.Atoi(*max)
		if err != nil {
			log.Panic(err)
		}
		if m < len(words) {
			maxWords = m
			wordsToGame = words[:m]
		}
	}
	log.Infof("will process %d words", maxWords)

	// start a new game for each word and count attempts
	for _, word := range wordsToGame {
		var win bool
		var attempts int
		var result string
		var err error
		progress++

		log.Infof("Try to guess word %s", word)
		lastWord := *firstWord
		g := wordle.New(words)

		// first attempt with firstWord ("taris")
		result, err = wordle.Try(word, lastWord, *upperCase)
		if err != nil {
			log.Panic(err)
		}
		log.Infof("\t[%d] guess: %s result: %s", attempts, green(lastWord), red(result))
		attempts++
		// Did we win with the first attempts?
		if result == resultOK {
			win = true
		}

		// next attempts
		for ; attempts < maxAttempts && !win; attempts++ {
			g.RemoveWord(lastWord)
			nextWord, _, err := g.NextWord(lastWord, result, *upperCase)
			if err != nil {
				log.Panic(err)
			}
			if nextWord == "" {
				break
			}
			result, err = wordle.Try(word, nextWord, *upperCase)
			if err != nil {
				log.Panic(err)
			}
			log.Infof("\t[%d] guess: %s result: %s", attempts, green(nextWord), red(result))
			if result == resultOK {
				win = true
			}
			lastWord = nextWord
		}
		if win {
			log.Infof("%s ✅ Found word %s in %d attempts - progress : %0.f %%", green("SUCCESS"), word, attempts, float64(progress)*100/float64(maxWords))
			successes = append(successes, attempts)
			results[attempts-1]++
		} else {
			log.Infof("%s ❌ Couldn't find word %s in %d attempts or less - progress : %0.f %%", red("FAILURE"), word, maxAttempts, float64(progress)*100/float64(maxWords))
			results[6]++
		}
	}
	var averageAttempts float64
	for _, a := range successes {
		averageAttempts += float64(a)
	}
	averageAttempts /= float64(len(successes))
	log.Infof("wordlebot performance is %f attempts to guess a word", averageAttempts)

	// create a new bar instance
	bar := charts.NewBar()
	// set some global options like Title/Legend/ToolTip or anything else
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Results",
		Subtitle: "Number of attempts",
	}))

	items := make([]opts.BarData, 0)
	for i := 0; i < 7; i++ {
		items = append(items, opts.BarData{Value: results[i]})
	}

	var xAxis []string
	for i := 1; i < 7; i++ {
		xAxis = append(xAxis, fmt.Sprintf("%d", i))
	}
	xAxis = append(xAxis, "failed")
	// Put data into instance
	bar.SetXAxis(xAxis).
		AddSeries("results", items)
	// Where the magic happens
	f, err := os.Create("results.html")
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()

	err = bar.Render(f)
	if err != nil {
		log.Panic(err)
	}

}
