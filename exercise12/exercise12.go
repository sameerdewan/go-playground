package exercise12

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func logRuntime() {
	fmt.Printf("\nos:%v\n", runtime.GOOS)
	fmt.Printf("\narch:%v\n", runtime.GOARCH)
	fmt.Printf("\ncpu(s):%v\n", runtime.NumCPU())
	fmt.Printf("\ngo routine(s):%v\n", runtime.NumGoroutine())
}

func Run() {
	logRuntime()

	foo := func() { // consistently slower over time
		for i := 0; i < 10; i++ {
			fmt.Printf("foo:%v\n", i)
			time.Sleep(time.Duration(int(time.Second) + i))
		}
		wg.Done()
	}

	bar := func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("bar:%v\n", i)
			if i > 5 { // fast at first, slows after half way by multiples
				time.Sleep(time.Duration(int(time.Second) * i / 2))
			}
		}
		wg.Done()
	}

	wg.Add(2)

	go foo()
	go bar()

	logRuntime()

	wg.Wait()
}
