package exercise7

import (
	"fmt"
	"math/rand"
)

type Location struct {
	latitude  float32
	longitude float32
}

type Human struct {
	age      uint
	first    string
	last     string
	location Location
	id       uint
}

func structsTest() {
	var location Location = Location{latitude: 87.3231, longitude: 99.1232}
	fmt.Printf("\nlocation:%v\n", location)
	var sameerDewan Human = Human{
		age:      27,
		first:    "Sameer",
		last:     "Dewan",
		location: location,
		id:       uint(rand.Int()),
	}
	fmt.Printf("\nhuman:%v\n", sameerDewan)
}

func demoMap() map[string]int {
	demo := map[string]int{
		"test":  1,
		"test2": 2,
	}
	return demo
}

func assignValue(demo map[string]int) {
	demo["test4"] = 4
}

func loopThroughMap(demo map[string]int) {
	fmt.Println("\nLooping through key values of demo...")
	for key, value := range demo {
		fmt.Printf("key:%v, value:%v\n", key, value) // maps do not have particular order
	}
}

func deleteOddValuedPairs(demo map[string]int) {
	fmt.Println("\nDeleting odd valued pairs...")
	for key, value := range demo {
		if value%2 != 0 {
			delete(demo, key)
		}
	}
}

func Run() {
	var demo map[string]int = demoMap()
	fmt.Printf("\n%v\n", demo)
	demo["test3"] = 3
	fmt.Printf("\n%v\n", demo)
	assignValue(demo) // even though this is passed to another function, demo IS mutated.
	fmt.Printf("\n%v\n", demo)
	loopThroughMap(demo)
	deleteOddValuedPairs(demo)
	fmt.Printf("\n%v\n", demo)
	structsTest()
}
