package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{25, -98, 155, -5, -6, 1, 1, 2, -1, 1}
	b := []int{0, 1, 2, 3, 4, 5}
	c := []int{9, 7, 4, 1, 3, 5}
	d := []int{0}
	s := []int{}
	z := []int{1, 1}

	fmt.Println(checkSliceIsSorted(a))
	fmt.Println(checkSliceIsSorted(b))
	fmt.Println(checkSliceIsSorted(c))
	fmt.Println(checkSliceIsSorted(d))
	fmt.Println(checkSliceIsSorted(s))
	fmt.Println(checkSliceIsSorted(z))
}

func checkSliceIsSorted(a []int) bool {
	sort.Ints(a)
	s := sort.IntsAreSorted(a)
	return s
}
