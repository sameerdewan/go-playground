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

func grabSlice(slice []interface{}, start uint, end uint) []interface{} {
	return slice[start:end]
}

func testAppend() {
	x := []int{1, 2, 3, 4, 5}
	test := append(x, 6)
	fmt.Printf("\nx:%v\ntest:%v", x, test) // x unmutated
	y := []int{6, 7, 8}
	x = append(x, y...)
	fmt.Printf("\nx:%v", x) // x mutated
}

func Run() {
	var slice []interface{} = generalizedSlice()
	logSliceIndexValue(slice)
	var slicedSlice []interface{} = grabSlice(slice, 0, 1)
	fmt.Printf("\nslice:%v\nslicedSlice:%v\n", slice, slicedSlice) // test shows slicing a slice does not mutate the original slice
	testAppend()
}
