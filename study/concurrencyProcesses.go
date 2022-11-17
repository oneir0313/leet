package study

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func runConcurrencyProcesses(concurrencyProcesses int, jobCount int) []int {

	var wg sync.WaitGroup
	wg.Add(jobCount)
	found := make(chan int)
	queue := make(chan int)

	go func() {
		for i := 0; i < jobCount; i++ {
			queue <- i
		}
		fmt.Println("queue close")

		close(queue)
	}()

	for i := 0; i < concurrencyProcesses; i++ {
		go func(queue <-chan int, found chan<- int) {
			for val := range queue {
				defer wg.Done()
				waitTime := rand.Int31n(1000)
				fmt.Println("job:", val, "wait time:", waitTime, "millisecond")
				time.Sleep(time.Duration(waitTime) * time.Millisecond)
				found <- val
			}
			fmt.Println("queue return")
		}(queue, found)
	}

	go func() {
		wg.Wait()
		close(found)
	}()
	var results []int
	for p := range found {
		fmt.Println("Finished job:", p)
		results = append(results, p)
	}

	fmt.Println("result:", results)
	return results
}
