package main

import (
	"fmt"
	q "m7/pkg"
	"strconv"
)

var firsNum, secNum, action string

func main() {

	fmt.Print("введите число a: ")
	fmt.Scan(&firsNum)
	fmt.Print("введите действие: ")
	fmt.Scan(&action)
	fmt.Print("введите число b: ")
	fmt.Scan(&secNum)
	aNum, err := strconv.ParseFloat(firsNum, 64)
	if err != nil {
		fmt.Println("Ошибка")
	}
	bNum, err := strconv.ParseFloat(secNum, 64)
	if err != nil {
		fmt.Println("Ошибка")
	}

	newCalc := q.Create() //создание экземпляра структуры calculate
	fmt.Println(newCalc.Calculate(aNum, bNum, action))

}
