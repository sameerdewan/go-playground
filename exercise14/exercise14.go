package exercise14

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func Run() {
	var counter int64

	var wg sync.WaitGroup
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("b - Current count:\t\t", atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Reached the end of the program b")

	fmt.Println("b - Final count:\t\t", atomic.LoadInt64(&counter))
}
