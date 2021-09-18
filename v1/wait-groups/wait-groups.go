//
// Adapted from https://goinbigdata.com/golang-wait-for-all-goroutines-to-finish/
// which seems to be very similar to
// https://gobyexample.com/waitgroups
//
// See also the example at https://golang.org/pkg/sync/
//

package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, id int) {
	defer wg.Done()

	fmt.Printf("Worker %v: Started\n", id)
	time.Sleep(time.Second / 10)
	fmt.Printf("Worker %v: Finished\n", id)
}

func toN(to int, f func(i int)) {
	for i := 1; i <= to; i++ {
		f(i)
	}
}

func main() {
	var wg sync.WaitGroup

	toN(4, func(i int) {
		fmt.Println("Main: Starting worker", i)
		wg.Add(1)
		go worker(&wg, i)
	})

	fmt.Println("Main: Waiting for workers to finish")
	wg.Wait()
	fmt.Println("Main: Completed")
}
