package main

import "fmt"

func main() {
	num := 3
	fmt.Println(countSheep(num))
}

func countSheep(num int) string {
	for i := 1; i <= num; i++ {
		switch {
		case i >= 1:
			return num + " sheep..."
		default:
			return "ошибка"
		}
	}
}

// if value == 0 {
// 	return ""
// }
// if value >= 1 {
// 	return value + " sheep..."
// }
