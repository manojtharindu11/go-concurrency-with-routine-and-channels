package examples

import "fmt"

func calculate(ch chan int) {
	result := 10 + 20
	ch <- result
}

func channelReceiving() {
	ch := make(chan int)

	go calculate(ch)

	result := <-ch

	fmt.Println("Result:", result)
}
