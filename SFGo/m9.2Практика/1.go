package main

import (
	"fmt"
	"math"
)

func main() {
	total := 1 + 1 + (-1)

	fmt.Println("Total : ", total)

	correctTotal := 1 + 1 + math.Abs(-1) // convert negative 1 to positive

	fmt.Println("Correct total with absolute : ", correctTotal)

}

func MakeNegative(x int) int {
	switch {
	case x == 0:
		return 0
	case x > 0:
		return x * -1
	case x < 0:
		return x
	}
	return 0
}
