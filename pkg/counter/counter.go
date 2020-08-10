package counter

// Counter хранит количество использований каждого символа
type Counter struct {
	statistic  map[rune]int
	symbolChan chan map[rune]int
	max        int
}

// Start запуск подсчета смиволов
func (c *Counter) Start() {
	go func() {
		for symbols := range c.symbolChan {
			for key, val := range symbols {
				c.statistic[key] += val
				if c.statistic[key] > c.max {
					c.max = c.statistic[key]
				}
			}
		}
	}()
}

// Stop остановка подсчета  символов
func (c *Counter) Stop() {
	close(c.symbolChan)
}

// Add сохраняет использование symbol
func (c *Counter) Add(symbols map[rune]int) {
	c.symbolChan <- symbols
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
		symbolChan: make(chan map[rune]int),
	}
}
