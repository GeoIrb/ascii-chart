package worker

import (
	"bufio"
	"io"
	"os"
	"sync"
)

type counter interface {
	Add(symbol map[rune]int)
}

// Worker ...
type Worker struct {
	counter counter
	end     <-chan struct{}
}

// Start of character counting in the file
func (w *Worker) Start(wg *sync.WaitGroup, fileName string) (err error) {
	defer func() {
		<-w.end
		wg.Done()
	}()

	var file *os.File
	if file, err = os.Open(fileName); err != nil {
		return
	}

	r := bufio.NewReader(file)
	var s rune
	symbols := make(map[rune]int)
	for err != io.EOF {
		s, _, err = r.ReadRune()
		symbols[s]++
	}
	w.counter.Add(symbols)
	return
}

// NewWorker ...
func NewWorker(
	counter counter,
	end chan struct{},
) *Worker {
	return &Worker{
		counter: counter,
		end:     end,
	}
}
