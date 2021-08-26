package exercise11

import (
	"fmt"
	"sort"
)

func sortedSliceExample() {
	slice := []string{"a", "b", "c"}
	fmt.Printf("\nslice:\t\t%v\t\t%v\t\t%v\n", slice[0], slice[1], slice[2])
	sort.Strings(slice)
	fmt.Println(slice)
	sort.Sort(sort.Reverse(sort.StringSlice(slice)))
	fmt.Println(slice)
}

func Run() {
	sortedSliceExample()
}
