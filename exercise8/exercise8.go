package exercise8

import "fmt"

func spreadParams(x ...int) { // "Variadic parameters"
	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}
}

func Run() {
	spreadParams(1, 2, 3, 4, 5)
	spreadable := []int{6, 7, 8, 9}
	spreadParams(spreadable...)
}
