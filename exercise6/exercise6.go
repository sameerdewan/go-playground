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

func Run() {
	var slice []interface{} = generalizedSlice()
	logSliceIndexValue(slice)
	var slicedSlice []interface{} = grabSlice(slice, 0, 1)
	fmt.Printf("\nslice:%v\nslicedSlice:%v\n", slice, slicedSlice) // test shows slicing a slice does not mutate the original slice
}
