package counter

import (
	"sync"
)

// Counter хранит количество использований каждого символа
type Counter struct {
	mutex     sync.Mutex
	max       int
	statistic map[rune]int
}

// Add сохраняет использование symbol
func (c *Counter) Add(symbol rune) {
	c.mutex.Lock()
	c.statistic[symbol]++
	if c.max < c.statistic[symbol] {
		c.max = c.statistic[symbol]
	}
	c.mutex.Unlock()
}

// GetValue вохращает количество использования symbol
func (c *Counter) GetValue(symbol rune) int {
	return c.statistic[symbol]
}

// GetMax возращает максимальное число испольщования символа
func (c *Counter) GetMax() int {
	return c.max
}

// NewCounter ...
func NewCounter() *Counter {
	return &Counter{
		statistic: make(map[rune]int),
	}
}
