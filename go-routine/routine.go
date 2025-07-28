package main

import (
	"fmt"
	"sync"
)

func printNum(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(i)
}

func main() {

	var wg sync.WaitGroup
	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go printNum(i, &wg)
	}
	wg.Wait()

}

// Go lang works on GMP model.
// G(Goroutines) - where Goroutines are lightweighted threads managed by the Go runtime.
// M(Machine) - OS threads on which Goroutines run
// P(Processor) - Scheduler that manages Goroutines

// Go uses M:N Scheduler, where Goroutines [G] are multiplexed over OS thread [M] manage by Processor[p] and
// his allows Go to run thousands of goroutines with very low memory and CPU overhead.
