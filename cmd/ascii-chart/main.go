package main

import (
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"sync"

	"github.com/GeoIrb/ascii-chart/pkg/counter"
	"github.com/GeoIrb/ascii-chart/pkg/graphic"
	"github.com/GeoIrb/ascii-chart/pkg/worker"
)

const (
	defaultDir = "/home/geo/go/src/github.com/GeoIrb/ascii-chart/test/data"

	barChartLayout = "\033[1;34m%s\033[0m\033[1;36m|\033[0m%s%d\n"
	barCharPart    = "\033[1;31m|\033[0m"

	firstSymbol = '!'
	lastSymbol  = '~'

	max    = 100
	border = 1
)

func main() {
	testDir := defaultDir
	if len(os.Args) > 1 {
		testDir = os.Args[1]
	}

	c5r := counter.NewCounter()
	c5r.Start()

	c := make(chan struct{}, runtime.NumCPU())
	w4r := worker.NewWorker(c5r, c)
	files, err := ioutil.ReadDir(testDir)
	if err != nil {
		log.Fatal(err)
	}

	wg := &sync.WaitGroup{}
	for _, file := range files {
		c <- struct{}{}
		wg.Add(1)
		go w4r.Start(wg, testDir+"/"+file.Name())
	}
	wg.Wait()
	c5r.Stop()

	g4c := graphic.NewGraphic(
		firstSymbol,
		lastSymbol,
		max,
		border,
		barChartLayout,
		barCharPart,
	)
	g4c.BarChart(c5r)
}
