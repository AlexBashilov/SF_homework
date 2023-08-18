package main

import "fmt"

func main() {
	arr := []int{-1, -2, -4, -5, -6}
	fmt.Println(findMax(arr))
}

func findMax(array []int) (max int, err error) {
	if len(array) == 0 {
		return 0, fmt.Errorf("could not found max in empty slice")
	}

	max = array[0]
	for _, val := range array[1:] {
		if val < max {
			max = val
		}
	}

	return max, nil
}
