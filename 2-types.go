package main

import "fmt"

func dynamicTypeTest() {
	var impliedInt interface{} = 5
	impliedInt = ""
	impliedInt = 10
	fmt.Println(impliedInt)
}

func main() {
	dynamicTypeTest()
}
