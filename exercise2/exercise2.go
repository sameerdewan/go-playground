package exercise2

import "fmt"

/*
	Exercise #2:
	------------
	Use var to DECLARE three VARIABLES.
	The variables should have package level scope. Do not assign VALUES to the variables.

	Use the following IDENTIFIERS for the variables and make sure the variables are of the following TYPE (meaning they can store VALUES of that TYPE).
		a. identifier “x” type int
		b. identifier “y” type string
		c. identifier “z” type bool
		d. print out the values for each identifier
		e. The compiler assigned values to the variables. What are these values called?
			Answer: Zero values
*/

var x int
var y string
var z bool

func Run() {
	fmt.Printf("x: %v\ny: %v\nz: %v\n", x, y, z)
}
