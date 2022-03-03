package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/jtbonhomme/wordlebot/internal/wordle"
	log "github.com/sirupsen/logrus"
)

func main() {
	var local = flag.String("l", "assets/taris.stat", "use local word stat file")
	var debug = flag.Bool("d", false, "display debug information")
	flag.Parse()
	if *local == "" {
		log.Panic("a valid word stat file is mandatory")
	}
	log.Infoln("start with local word stat file", *local)

	if *debug {
		log.SetLevel(log.DebugLevel)
	}

	statFile, err := os.Open(*local)
	if err != nil {
		log.Panic(err)
	}
	defer statFile.Close()
	reader := csv.NewReader(statFile)
	statData, err := reader.ReadAll()
	if err != nil {
		log.Panic(err)
	}

	items := make([]opts.BarData, 0)

	for _, each := range statData {
		e, err := strconv.ParseFloat(each[1], 64)
		if err != nil {
			log.Panic(err)
		}

		items = append(items, opts.BarData{Value: e})
	}

	// create a new bar instance
	bar := charts.NewBar()
	// set some global options like Title/Legend/ToolTip or anything else
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Entropy repartition for word " + *local,
		Subtitle: "X axis represents the information gained for this word",
	}))

	var xAxis []string
	for i := 0; i < int(math.Pow(3, 5)); i++ {
		result := wordle.IntToPowerOf3(i)
		xAxis = append(xAxis, fmt.Sprintf("%d%d%d%d%d", result[0], result[1], result[2], result[3], result[4]))
	}
	// Put data into instance
	bar.SetXAxis(xAxis).
		AddSeries("Entropy", items)
	// Where the magic happens
	f, err := os.Create("bar.html")
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()

	err = bar.Render(f)
	if err != nil {
		log.Panic(err)
	}
}
