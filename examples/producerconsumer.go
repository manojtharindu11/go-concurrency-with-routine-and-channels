package examples

import (
	"fmt"
	"sync"
)

func producer(jobs chan<- int) {
	defer close(jobs)

	for i := 1; i <= 5; i++ {
		fmt.Println("Producing jobs:", i)
		jobs <- i
	}
}

func consumer(jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Println("Consuming job:", job)
	}

	fmt.Println("Consumer finished")
}

func JobHandler() {
	jobs := make(chan int, 2)

	var wg sync.WaitGroup
	wg.Add(1)

	go consumer(jobs, &wg)

	producer(jobs)

	wg.Wait()

	fmt.Println("Program finished")
}
