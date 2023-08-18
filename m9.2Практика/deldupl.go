package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MergeArrays([]int{1, 2, 3, 4}, []int{5, 6, 7, 8}))
}

// Данная функция совмещает срезы
// и сортирует их по возрастанию
func MergeArrays(arr1, arr2 []int) []int {
	res := []int{}
	res = append(arr1, arr2...)
	sort.Ints(res)
	return DelDupl(res)
}

// Функция принимает массив из функции MergeArrays
// и удаляет дубликаты в срезе
func DelDupl(res []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range res {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
