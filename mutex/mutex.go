package main

import (
	"fmt"
	"sync"
)

type social struct {
	views int
	mu    sync.Mutex
}

func (s *social) inc(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		s.mu.Unlock()
	}()

	s.mu.Lock()
	s.views += 1
}

func main() {
	var wg sync.WaitGroup

	ig := social{views: 0}

	for i := 0; i < 40; i++ {
		wg.Add(1)
		go ig.inc(&wg)

	}

	wg.Wait()

	fmt.Println(ig.views)
}
