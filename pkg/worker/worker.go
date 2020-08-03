package worker

import (
	"bufio"
	"io"
	"os"
	"sync"
)

type counter interface {
	Add(symbol rune)
}

type Worker struct {
	counter counter
}

func (w *Worker) Start(wg *sync.WaitGroup, fileName string) (err error) {
	defer wg.Done()

	var file *os.File
	if file, err = os.Open(fileName); err != nil {
		return
	}
	r := bufio.NewReader(file)
	var s rune
	for err != io.EOF {
		s, _, err = r.ReadRune()
		w.counter.Add(s)
	}
	return
}

func NewWorker(
	counter counter,
) *Worker {
	return &Worker{
		counter: counter,
	}
}
