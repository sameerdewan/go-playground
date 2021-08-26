package exercise1

import "fmt"

/*
	Exercise #1:
	------------
	Using the short declaration operator, ASSIGN these VALUES to VARIABLES with the IDENTIFIERS “x” and “y” and “z”
		42
		“James Bond”
		true
	Now print the values stored in those variables using
		a. single print statement
		b. multiple print statements
*/

func Run() {
	x, y, z := 42, "James Bond", true
	fmt.Println("a\n----------------------------------")
	fmt.Printf("x: %v\ny: %v\nz: %v\n", x, y, z) // a
	fmt.Println("-----------------------------------")
	fmt.Println("b\n----------------------------------")
	fmt.Printf("x: %v\n", x) // b.1
	fmt.Printf("y: %v\n", y) // b.2
	fmt.Printf("z: %v\n", z) // b.3
}
