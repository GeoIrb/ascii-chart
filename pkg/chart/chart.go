package chart

import (
	"fmt"
)

type counter interface {
	GetCount(symbol string) int
	GetMax() int
}

// Chart data for output chart
type Chart struct {
	border   *int
	maxXAxis int
	maxYAxis int

	layout string
	part   string
}

// Bar output bar chart
func (c *Chart) Bar(data counter) {
	for i := 33; i < c.maxXAxis; i++ {
		if count := data.GetCount(string(i)); c.border != nil && *c.border < count {
			fmt.Printf(c.layout, string(i), c.line(count, data.GetMax()), count)
		}
	}
}

func (c *Chart) line(size, max int) (line string) {
	for i := 0; i < int(float64(size)/float64(max)*float64(c.maxYAxis)); i++ {
		line += c.part
	}
	return
}

func NewChart(
	border *int,
	maxXAxis int,
	maxYAxis int,

	layout string,
	part string,
) *Chart {
	return &Chart{
		border:   border,
		maxXAxis: maxXAxis,
		maxYAxis: maxYAxis,

		layout: layout,
		part:   part,
	}
}
