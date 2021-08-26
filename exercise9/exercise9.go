package exercise9

import "fmt"

func mutateBy2x(number *int) {
	fmt.Printf("\nPre-mutate >\t\tnumber:%v\taddress:%v\n", *number, &number)
	*number = *number * 2
	fmt.Printf("\nPost-mutate >\t\tnumber:%v\taddress:%v\n", *number, &number)
}

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

	d := 2
	mutateBy2x(&d) // same pointer in memory but new value
}
