package exercise17

import "fmt"

func Run() {
	even := make(chan int)
	odd := make(chan int)
	end := make(chan string)

	// send
	go send(even, odd, end)

	// receive
	receive(even, odd, end)

	fmt.Println("\nExiting 17")
}

func send(even, odd chan<- int, end chan<- string) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	end <- "This is the end!"
}

func receive(even, odd <-chan int, end <-chan string) {
	for {
		select {
		case num := <-even:
			fmt.Printf("\nEven: %v", num)
		case num := <-odd:
			fmt.Printf("\nOdd: %v", num)
		case msg := <-end:
			fmt.Println("\nEnd: ", msg)
			return
		}
	}
}
