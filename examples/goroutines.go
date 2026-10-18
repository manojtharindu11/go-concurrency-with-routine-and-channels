package examples

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello from the function")
}

func routineExample() {
	go sayHello()

	time.Sleep(time.Second)

	fmt.Println("Main function finished")
}
