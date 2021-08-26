package exercise6

import "fmt"

func generalizedSlice() []interface{} {
	test := []interface{}{1, "this is a string value", true, 6.4} // generalized slice
	fmt.Println(test)
	return test
}

func logSliceIndexValue(slice []interface{}) {
	for index, value := range slice {
		fmt.Printf("\nindex: %v\t\tvalue: %v\n", index, value)
	}
}

func Run() {
	var slice []interface{} = generalizedSlice()
	logSliceIndexValue(slice)
}
