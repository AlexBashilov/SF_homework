package main

import (
	"fmt"
)

var A, B float32
var Math string

func main() {
	Ascan()
	Action()
	Bscan()

	switch {
	case Math == "+":
		fmt.Println(A + B)
	case Math == "-":
		fmt.Println(A - B)
	case Math == "/":
		fmt.Println(A / B)
	case Math == "*":
		fmt.Println(A * B)
	default:
		fmt.Println("Действие не найдено. Введите по новой!")
	}
}

func Ascan() {
	fmt.Print("введите число a: ")
	_, err := fmt.Scanln(&A)
	if err != nil {
		fmt.Println(fmt.Sprintf("не получилось считать число: %s", err))
	}
}

func Bscan() {
	fmt.Print("введите число b: ")
	_, err := fmt.Scanln(&B)
	if err != nil {
		fmt.Println(fmt.Sprintf("не получилось считать число: %s", err))
	}
	if B == 0 {
		fmt.Println("На ноль делить нельзя!!")
	}
}

func Action() {
	fmt.Print("введите действие: ")
	_, err := fmt.Scanln(&Math)
	if err != nil {
		fmt.Println(fmt.Sprintf("не получилось считать действие: %s", err))
	}
}
