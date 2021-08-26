package exercise7

import "fmt"

func demoMap() map[string]int {
	demo := map[string]int{
		"test":  1,
		"test2": 2,
	}
	return demo
}

func assignValue(demo map[string]int) {
	demo["test4"] = 4
}

func Run() {
	var demo map[string]int = demoMap()
	fmt.Printf("\n%v\n", demo)
	demo["test3"] = 3
	fmt.Printf("\n%v\n", demo)
	assignValue(demo) // even though this is passed to another function, demo IS mutated.
	fmt.Printf("\n%v\n", demo)
}
