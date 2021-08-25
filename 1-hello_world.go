package main

import (
	"fmt"
	"os"
	"text/template"

	"github.com/sameerdewan/go-playground/importable"
)

type Child struct {
	Name string
	Age  int
}

type Parent struct {
	Name  string
	Child string
}

func congradulateParents(parents []string, child string) {
	for i := 0; i < len(parents); i++ {
		msg := "Congradulations, {{.Name}} on the birth of your child, {{.Child}}!\n"
		substitute := Parent{Name: parents[i], Child: child}
		message, error := template.New("msg").Parse(msg)
		message.Execute(os.Stdout, substitute)
		if error != nil {
			fmt.Println(error)
		}
	}
}
func growAYearOlder(name string, age int) {
	msg := "Happy Birthday, {{.Name}}. You are now {{.Age}} year(s) old.\n"
	substitute := Child{Name: name, Age: age}
	message, error := template.New("msg").Parse(msg)
	message.Execute(os.Stdout, substitute)
	if error != nil {
		fmt.Println(error)
	}
}

func dynamicTypeTest() int {
	var impliedInt interface{} = 5
	impliedInt = ""
	impliedInt = 10
	fmt.Println(impliedInt)
	coercedInt := impliedInt.(int)
	return coercedInt
}

func main() {
	year := 27
	parents := []string{"Thomas", "Sarah"}
	child := "Thomas, Jr."
	for i := 0; i <= year; i++ {
		if i == 0 {
			congradulateParents(parents, child)
		} else {
			growAYearOlder(child, i)
		}
	}
	var integer int = dynamicTypeTest()
	fmt.Printf("%T", integer) // Prints type
	importable.RunImportable()
}
