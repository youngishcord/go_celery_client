package main

import (
	"fmt"
	"sync"
	"time"
)

type Worker struct {
	name string
}

func (w *Worker) Add(x, y int) int {
	return x + y
}

func main() {
	w := Worker{
		name: "test",
	}

	var wg sync.WaitGroup

	for g := 1; g < 4; g++ {

		wg.Add(1)
		go func(g int) {
			defer wg.Done()
			name := fmt.Sprintf("gorutine %d", g)
			for i := 0; i < 10; i++ {
				fmt.Println(name, ": ", w.Add(i, i+1))
				time.Sleep(time.Second * time.Duration(g))
			}
		}(g)
	}

	wg.Wait()

}
