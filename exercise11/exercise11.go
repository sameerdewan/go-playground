package exercise11

import (
	"fmt"
	"sort"

	"golang.org/x/crypto/bcrypt"
)

func sortedSliceExample() {
	slice := []string{"a", "b", "c"}
	fmt.Printf("\nslice:\t\t%v\t\t%v\t\t%v\n", slice[0], slice[1], slice[2])
	sort.Strings(slice)
	fmt.Println(slice)
	sort.Sort(sort.Reverse(sort.StringSlice(slice)))
	fmt.Println(slice)
}

func exampleBcrypt() {
	rawPassword := "password123"
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\nraw:%v\t\tencrypted:%v\n", rawPassword, encryptedPassword)
}

func Run() {
	sortedSliceExample()
	exampleBcrypt()
}
