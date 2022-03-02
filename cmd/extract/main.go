package main

import (
	"encoding/csv"
	"flag"
	"os"
	"regexp"
	"sort"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// Word is a struct that records all french lemmes and their frequancy in movies
type Word struct {
	Lemme     string
	Frequency float64
}

// ByFrequency implements sort.Interface for []Word based on
// the Frequency field.
type ByFrequency []Word

func (a ByFrequency) Len() int           { return len(a) }
func (a ByFrequency) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByFrequency) Less(i, j int) bool { return a[i].Frequency < a[j].Frequency }

func main() {
	var validLemme = regexp.MustCompile(`^[a-z]{5}$`)
	var local = flag.String("l", "assets/Lexique383.tsv", "use local lexical base")
	var debug = flag.Bool("d", false, "display debug information")
	flag.Parse()
	log.Infoln("start with local lexical database ", *local)

	if *debug {
		log.SetLevel(log.DebugLevel)
	}

	// read data from TSV file
	tsvFile, err := os.Open(*local)
	if err != nil {
		log.Panic(err)
	}
	defer tsvFile.Close()
	reader := csv.NewReader(tsvFile)
	reader.Comma = '\t' // Use tab-delimited instead of comma
	reader.FieldsPerRecord = -1
	tsvData, err := reader.ReadAll()
	if err != nil {
		log.Panic(err)
	}

	// Parse TSV
	var records []Word
	check := make(map[string]int)

	log.Debugln("lemme freq")
	var l int
	for i, each := range tsvData {
		lemme := each[0]
		freqString := each[7]
		// remove header line
		if i == 0 {
			continue
		}
		// remove invalid lemmes
		if !validLemme.MatchString(lemme) {
			continue
		}
		// remove duplicates
		_, ok := check[lemme]
		if ok {
			continue
		}

		check[lemme] = 1
		log.Debugf("%d %s %s", l, lemme, freqString)
		l++
		freq, err := strconv.ParseFloat(freqString, 64)
		if err != nil {
			log.Panic(err)
		}

		record := Word{
			Lemme:     lemme,
			Frequency: freq,
		}
		records = append(records, record)
	}

	// Sort records
	sort.Sort(ByFrequency(records))
	log.Infof("total records %d", len(records))

	sortedLemmes, err := os.Create("assets/words.txt")
	if err != nil {
		log.Panic(err)
	}
	defer sortedLemmes.Close()

	for _, record := range records {
		_, err := sortedLemmes.Write([]byte(record.Lemme + "\n"))
		if err != nil {
			log.Panic(err)
		}
	}
}
