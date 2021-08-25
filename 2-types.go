package main

import "fmt"

func dynamicTypeTest() int {
	var impliedInt interface{} = 5
	impliedInt = ""
	impliedInt = 10
	fmt.Println(impliedInt)
	coercedInt := impliedInt.(int)
	return coercedInt
}

func main() {
	var integer int = dynamicTypeTest()
	fmt.Printf("%T", integer)
}
