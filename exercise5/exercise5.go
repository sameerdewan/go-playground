package exercise5

import "fmt"

func printNumAsDecBinHex(number int) {
	var intToFloat float64 = float64(number)
	fmt.Printf("\ndecimal:%f\nbinary:%b\nhexidecimal:%x\n\n", intToFloat, number, number)
}

const (
	example1 = iota
	example2 = iota
	example3 = iota
)

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

	printNumAsDecBinHex(3)

	exampleStringLiteral := `
		This string literal will retain formatting

			ie: line skipping 


		ie:          spacing,
	`

	fmt.Println(exampleStringLiteral)

	fmt.Printf("\n\nexample1:%v\nexample2:%v\nexample3:%v\n", example1, example2, example3)
}
