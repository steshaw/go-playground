//
// Adapted from https://goinbigdata.com/golang-wait-for-all-goroutines-to-finish/
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

func main() {
	var wg sync.WaitGroup

	for i, _ := range [3]int{} {
		fmt.Println("Main: Starting worker", i)
		wg.Add(1)
		go worker(&wg, i)
	}

	fmt.Println("Main: Waiting for workers to finish")
	wg.Wait()
	fmt.Println("Main: Completed")
}
