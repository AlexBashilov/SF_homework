package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func main() {
	ar := make([]int, 50)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайное значение от [-100;100]
	}

	//selectionSort(ar)
	selectionSortByMax(ar)

	fmt.Println(ar)
}

// сортировка выбором слево направо
func selectionSort(ar []int) {
	for i := 0; i < len(ar); i++ {
		var minIndex = i
		for j := i + 1; j < len(ar); j++ {
			if ar[j] < ar[minIndex] {
				minIndex = j
			}
		}

		ar[i], ar[minIndex] = ar[minIndex], ar[i]
	}
}

// сортировка выбором справа налево
func selectionSortByMax(ar []int) {
	for i := len(ar) - 1; i >= 0; i-- {
		var maxIndex = i
		for j := i - 1; j >= 0; j-- {
			if ar[j] > ar[maxIndex] {
				maxIndex = j
			}
		}

		ar[i], ar[maxIndex] = ar[maxIndex], ar[i]
	}
}


// двунаправленная сортировка выбором
func selectionSortBidirectional(ar []int) {
	maxPos := len(ar) - 1
	minPos := 0

	for minPos <= maxPos {
		var minIdx = minPos
		var maxIdx = maxPos
		for j := minPos; j <= maxPos; j++ {
			if ar[j] < ar[minIdx] {
				minIdx = j
			}
			if ar[j] > ar[maxIdx] {
				maxIdx = j
			}
		}

		ar[minPos], ar[minIdx] = ar[minIdx], ar[minPos]
		if maxIdx == minPos {
			maxIdx = minIdx
		}
		ar[maxPos], ar[maxIdx] = ar[maxIdx], ar[maxPos]

		minPos++
		maxPos--
	}
}
