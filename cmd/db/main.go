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
		log.Debugf("%d %s %s", l, lemme, each[7])
		l++
		freq, err := strconv.ParseFloat(each[7], 64)
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
	log.Debugf("%v", records[0])
	log.Debugf("%v", records[len(records)-1])
	log.Debugf("total records %d", len(records))
	//if len(records) < 4096 {
	//	log.Panic("not enough records")
	//}
	//	mostFrequentLemmes := records[len(records)-4096:]
	mostFrequentFile, err := os.Create("assets/words.txt")
	if err != nil {
		log.Panic(err)
	}
	defer mostFrequentFile.Close()

	for _, record := range records {
		_, err := mostFrequentFile.Write([]byte(record.Lemme + "\n"))
		if err != nil {
			log.Panic(err)
		}
	}
	/*




		jsondata, err := json.Marshal(records) // convert to JSON

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// sanity check
		// NOTE : You can stream the JSON data to http service as well instead of saving to file
		fmt.Println(string(jsondata))

		// now write to JSON file

		jsonFile, err := os.Create("./data.json")

		if err != nil {
			fmt.Println(err)
		}

		var record Employee
		var records []Employee

		for _, each := range csvData {
			record.Name = each[0]
			record.Age, _ = strconv.Atoi(each[1]) // need to cast integer to string
			record.Job = each[2]
			records = append(records, record)
		}

		jsondata, err := json.Marshal(records) // convert to JSON

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// sanity check
		// NOTE : You can stream the JSON data to http service as well instead of saving to file
		fmt.Println(string(jsondata))

		// now write to JSON file

		jsonFile, err := os.Create("./data.json")

		if err != nil {
			fmt.Println(err)
		}
		defer jsonFile.Close()

		jsonFile.Write(jsondata)
		jsonFile.Close()
	*/
}
