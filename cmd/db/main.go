package main

import (
	"encoding/csv"
	"flag"
	"os"
	"regexp"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type Word struct {
	Lemme     string
	Frequency float64
}

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
	for i, each := range tsvData {
		// remove header line
		if i == 0 {
			continue
		}
		// remove invalid lemmes
		if !validLemme.MatchString(each[2]) {
			continue
		}
		// remove duplicates
		_, ok := check[each[2]]
		if ok {
			continue
		}

		check[each[2]] = 1
		log.Debugf("%s %s", each[2], each[7])
		freq, err := strconv.ParseFloat(each[7], 64)
		if err != nil {
			log.Panic(err)
		}

		record := Word{
			Lemme:     each[2],
			Frequency: freq,
		}
		records = append(records, record)
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
