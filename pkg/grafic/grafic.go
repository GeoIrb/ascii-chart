package grafic

import (
	"fmt"
)

type function interface {
	GetValue(symbol rune) int
	GetMax() int
}

// Grafic ...
type Grafic struct {
	minX int
	maxX int
	minY int
	maxY int

	layout string
	part   string
}

// BarChart  вывод гистограммы
func (c *Grafic) BarChart(f function) {
	max := f.GetMax()
	for i := 33; i < c.maxX; i++ {
		if count := f.GetValue(rune(i)); c.minY < count {
			scale := c.scale(count, max)
			fmt.Printf(c.layout, string(i), c.line(scale), count)
		}
	}
}

func (c *Grafic) line(length int) (line string) {
	for i := 0; i < length; i++ {
		line += c.part
	}
	return
}

func (c *Grafic) scale(size, max int) int {
	return int(float64(size) / float64(max) * float64(c.maxY))
}

// NewGrafic ...
func NewGrafic(
	minX int,
	maxX int,
	minY int,
	maxY int,

	layout string,
	part string,
) *Grafic {
	return &Grafic{}
}
