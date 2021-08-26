package exercise9

import "fmt"

func Run() {
	a := 17
	b := a // different pointers
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(b)
	fmt.Println(&b)
}
