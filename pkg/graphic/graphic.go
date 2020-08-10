package graphic

import (
	"fmt"
)

type function interface {
	GetValue(symbol rune) int
	GetMax() int
}

// Graphic ...
type Graphic struct {
	minX int
	maxX int
	maxY int

	border int

	layout string
	part   string
}

// BarChart  вывод гистограммы
func (g *Graphic) BarChart(f function) {
	max := f.GetMax()
	for i := 33; i < g.maxX; i++ {
		if count := f.GetValue(rune(i)); g.border <= count {
			scale := g.scale(count, max)
			fmt.Printf(g.layout, string(i), g.line(scale), count)
		}
	}
}

func (g *Graphic) line(length int) (line string) {
	for i := 0; i < length; i++ {
		line += g.part
	}
	return
}

func (g *Graphic) scale(size, max int) int {
	return int(float64(size) / float64(max) * float64(g.maxY))
}

// NewGraphic ...
func NewGraphic(
	minX int,
	maxX int,
	maxY int,

	border int,

	layout string,
	part string,
) *Graphic {
	return &Graphic{
		minX: minX,
		maxX: maxX,
		maxY: maxY,

		border: border,

		layout: layout,
		part:   part,
	}
}
