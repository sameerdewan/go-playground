package exercise3

import "fmt"

/*
	Exercise #3:
	------------
	Create your own type based off of a type at the global scope and then use conversion to
	convert the type of the value to the underlying type
*/

type melon string

var watermelon melon
var cantelope string

func Run() {
	fmt.Printf("type of watermelon: %T\n", watermelon)
	fmt.Printf("type of cantelope: %T\n", cantelope)
	cantelope = string(watermelon) // should compile
	watermelon = melon(cantelope)  // should compile
}
