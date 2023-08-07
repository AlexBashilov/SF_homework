package main

import "fmt"

func main() {
	var a1, a2, b1, b2, c1, c2, d1, d2 float64
	a1 = 10
	a2 = 75
	b1 = 69
	b2 = 23
	c1 = a1 - a2
	c2 = a2 - a1
	d1 = b1 - b2
	d2 = b2 - b1
	if c1/c2 == d1/d2 {
		fmt.Println("Прямые пересекаются")
	}
	if c1/c2 != d1/d2 {
		fmt.Println("Прямые НЕ пересекаются")

	}
}
