package exercise5

import "fmt"

func Run() {
	a := 4
	fmt.Printf("a) %d\t\t%b\n", a, a)

	b := a << 1
	fmt.Printf("b) %d\t\t%b\n", b, b)

	c := a | b
	fmt.Printf("c) %d\t\t%b\n", c, c)

	d := 8
	e := a | d
	fmt.Printf("e) %d\t\t%b\n", e, e)

	f := 18
	fmt.Printf("f) %d\t\t%b\n", f, f)
	g := a | f
	fmt.Printf("g) %d\t\t%b\n", g, g)
}
