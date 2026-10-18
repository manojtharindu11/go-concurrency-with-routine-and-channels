package examples

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		result := job * job

		fmt.Printf("Worker %d processed job %d\n", id, job)

		results <- result
	}
}

func WorkerWorks() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	var wg sync.WaitGroup

	// Start two workers.
	wg.Add(2)

	go worker(1, jobs, results, &wg)
	go worker(2, jobs, results, &wg)

	// Send five jobs.
	for i := 1; i <= 5; i++ {
		jobs <- i
	}

	// Tell workers that no more jobs are coming.
	close(jobs)

	// Wait until both workers have finished.
	wg.Wait()

	// No workers remain to send results.
	close(results)

	// Collect and print every result.
	for result := range results {
		fmt.Println("Result:", result)
	}

	fmt.Println("All jobs completed!")
}
