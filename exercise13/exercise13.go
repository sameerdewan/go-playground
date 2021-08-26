package exercise13

import (
	"fmt"
	"sync"
)

func Run() {
	var counter int64

	var wg sync.WaitGroup
	wg.Add(100)

	var mu sync.Mutex

	for i := 0; i < 100; i++ {
		go func() {
			mu.Lock()
			counter++
			fmt.Println("a - Current count:\t\t", counter)
			mu.Unlock()
			wg.Done()
		}()
	}

	fmt.Println("Reached the end of the program a")

	wg.Wait()

	fmt.Println("a - Final count:\t\t", counter)
}
