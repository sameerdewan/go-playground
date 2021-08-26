package exercise6

import "fmt"

func generalizedSlice() {
	test := []interface{}{1, "this is a string value", true, 6.4} // generalized slice
	fmt.Println(test)
}

func Run() {
	generalizedSlice()
}
