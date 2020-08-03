package main

import (
	"io/ioutil"
	"log"
	"os"
	"sync"

	"github.com/GeoIrb/ascii-chart/pkg/counter"
	"github.com/GeoIrb/ascii-chart/pkg/grafic"
	"github.com/GeoIrb/ascii-chart/pkg/worker"
)

const (
	defaultDir = "/home/geo/go/src/github.com/GeoIrb/ascii-chart/test/data"

	barChartLayout = "\033[1;34m%s\033[0m\033[1;36m|\033[0m%s%d\n"
	barCharPart    = "\033[1;31m|\033[0m"

	firstSymbol = '!'
	lastSymbol  = '~'

	max    = 150
	border = 1
)

func main() {
	testDir := defaultDir

	if len(os.Args) > 1 {
		testDir = os.Args[1]
	}

	files, err := ioutil.ReadDir(testDir)
	if err != nil {
		log.Fatal(err)
	}

	c5r := counter.NewCounter()
	go c5r.Start()
	defer c5r.Stop()
	w4r := worker.NewWorker(
		c5r,
	)

	wg := &sync.WaitGroup{}
	for _, file := range files {
		wg.Add(1)
		go w4r.Start(wg, testDir+"/"+file.Name())
	}
	wg.Wait()

	g4c := grafic.NewGrafic(
		firstSymbol,
		lastSymbol,
		max,
		border,
		barChartLayout,
		barCharPart,
	)
	g4c.BarChart(c5r)
}
