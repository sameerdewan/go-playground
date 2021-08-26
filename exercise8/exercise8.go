package exercise8

import "fmt"

type Pet interface {
	greetHuman()
}
type Dog string
type Cat string

func (dog Dog) bark() {
	fmt.Printf("Woof! I am a %v\n", dog)
}

func (dog Dog) greetHuman() {
	fmt.Printf("\nWoof! I am a %v\n", dog)
}

func (cat Cat) greetHuman() {
	fmt.Printf("\nMweow! I am a %v\n", cat)
}

func wagTail(pet Pet) {
	fmt.Printf("\n*%T wags tail*\n", pet)
}

func spreadParams(x ...int) { // "Variadic parameters"
	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}
}

func demoAnonymousFunc() {
	var germanShepherd Dog = "German Shepherd"
	func(dog Dog) {
		fmt.Printf("\n%v\n", germanShepherd)
	}(germanShepherd)
}

func returnFunc() func(aString string) {
	return func(aString string) {
		fmt.Printf("\nreturnFunc: %v\n", aString)
	}
}

func Run() {
	spreadParams(1, 2, 3, 4, 5)
	spreadable := []int{6, 7, 8, 9}
	spreadParams(spreadable...)
	var goldenRetriever Dog = "Golden Retriever"
	goldenRetriever.bark()
	var shibaInu Dog = "Shiba Inu"
	shibaInu.bark()

	var shihTzu Dog = "Shih Tzu"
	var havanaBrown Cat = "Havana Brown"

	shihTzu.greetHuman()
	havanaBrown.greetHuman()

	wagTail(shihTzu)
	wagTail(havanaBrown)

	demoAnonymousFunc()

	storedFunction := func() {
		fmt.Println("This was stored!")
	}

	storedFunction()

	printString := returnFunc()
	printString("My String")
}
