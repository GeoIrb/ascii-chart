package counter

import (
	"sync"
)

// Counter stores the number of used characters
type Counter struct {
	mutex     sync.Mutex
	max       int
	statistic map[string]int
}

// Add preserves symbol usage
func (c *Counter) Add(symbol string) {
	c.mutex.Lock()
	c.statistic[symbol]++
	if c.max < c.statistic[symbol] {
		c.max = c.statistic[symbol]
	}
	c.mutex.Unlock()
}

// GetCount number of used characters
func (c *Counter) GetCount(symbol string) int {
	return c.statistic[symbol]
}

// GetMax number of used characters
func (c *Counter) GetMax() int {
	return c.max
}

// NewCounter ...
func NewCounter() *Counter {
	return &Counter{
		statistic: make(map[string]int),
	}
}
