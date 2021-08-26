package exercise16

import "fmt"

func Run() {
	channel := make(chan int)

	// send
	go func() {
		for i := 1; i <= 100; i++ {
			channel <- i
		}
		close(channel)
	}()

	// receive
	for v := range channel {
		fmt.Println(v)
	}

	fmt.Println("End")
}
