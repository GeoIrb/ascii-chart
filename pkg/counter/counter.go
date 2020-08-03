package counter

// Counter хранит количество использований каждого символа
type Counter struct {
	statistic  map[rune]int
	symbolChan chan rune
	max        int
}

// Start запуск подсчета смиволов
func (c *Counter) Start() {
	for symbol := range c.symbolChan {
		c.statistic[symbol]++
		if c.statistic[symbol] > c.max {
			c.max = c.statistic[symbol]
		}
	}
}

// Stop остановка подсчета  символов
func (c *Counter) Stop() {
	close(c.symbolChan)
}

// Add сохраняет использование symbol
func (c *Counter) Add(symbol rune) {
	c.symbolChan <- symbol
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
		statistic:  make(map[rune]int),
		symbolChan: make(chan rune),
	}
}
