package exercise18

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func demoIo() {
	f, err := os.Create("./exercise18/test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	r := strings.NewReader("This was written into the file using go!")

	io.Copy(f, r)
}

func demoIo2() {
	f, err := os.Open("./exercise18/test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(bs)
	fmt.Println(string(bs))
}

func Run() {
	MyError := errors.New("MyError")
	fmt.Printf("type: %t\t\tlogs out -> %v\n", MyError, MyError)

	n, err := fmt.Println("Hello World!")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(n)

	demoIo()

	demoIo2()
}
