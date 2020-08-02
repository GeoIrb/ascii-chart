package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/GeoIrb/ascii-chart/pkg/counter"
	"github.com/GeoIrb/ascii-chart/pkg/grafic"
)

const (
	defaultDir = "/home/geo/go/src/github.com/GeoIrb/ascii-chart/test/data"

	barChartLayout = "\033[1;34m%s\033[0m\033[1;36m|\033[0m%s%d\n"
	barCharPart    = "\033[1;31m|\033[0m"
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

	fileName := testDir + "/" + files[1].Name()
	f, _ := os.Open(fileName)
	r := bufio.NewReader(f)

	c5r := counter.NewCounter()
	var s rune
	for err != io.EOF {
		s, _, err = r.ReadRune()
		c5r.Add(s)
	}

	g4c := grafic.NewGrafic(
		33,
		256,
		0,
		150,
		barChartLayout,
		barCharPart,
	)
	g4c.BarChart(c5r)

}
