package exercise9

import "fmt"

func Run() {
	a := 17
	b := a
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(b)
	fmt.Println(&b)

	c := &a
	fmt.Println(c)
	fmt.Println(*c)
	fmt.Println(&a)

	*c = 27
	fmt.Println(a)
}
