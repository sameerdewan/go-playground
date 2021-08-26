package exercise8

import "fmt"

type Dog string

func (dog Dog) bark() {
	fmt.Printf("Woof! I am a %v\n", dog)
}

func spreadParams(x ...int) { // "Variadic parameters"
	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}
}

func Run() {
	spreadParams(1, 2, 3, 4, 5)
	spreadable := []int{6, 7, 8, 9}
	spreadParams(spreadable...)
	var goldenRetriever Dog = "Golden Retriever"
	goldenRetriever.bark()
	var shibaInu Dog = "Shiba Inu"
	shibaInu.bark()
}
