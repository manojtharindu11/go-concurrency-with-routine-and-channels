package examples

import (
	"fmt"
	"sync"
)

func doTask(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Task %d started\n", id)

	fmt.Printf("Task %d finished\n", id)
}

func WaitGroupTask() {
	var wg sync.WaitGroup

	wg.Add(3)

	go doTask(1, &wg)
	go doTask(2, &wg)
	go doTask(3, &wg)

	wg.Wait()

	fmt.Println("All tasks completed!")
}
