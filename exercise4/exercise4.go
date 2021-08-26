package exercise4

import "fmt"

func Run() {
	x, y := 1.3, 2.2
	var sum int = int(x + y)
	var example string = fmt.Sprintf("Sum: %v", sum)
	fmt.Printf("The example return value is ---> %v\n", example)
}
