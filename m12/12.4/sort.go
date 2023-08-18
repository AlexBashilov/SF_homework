package sort

import (
	"math/rand"
)


//сортировка пузырьком
func bubbleSort(ar []int) {
	for i := 0; i < len(ar)-1; i++ {
		for j := 1; j < len(ar)-i; j++ {
			if ar[j-1] > ar[j] {
				ar[j-1], ar[j] = ar[j], ar[j-1]
			}
		}
	}
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


// сортировка слиянием
func mergeSort(ar []int) []int {
	if len(ar) < 2 {
		return ar
	}

	middle := len(ar) / 2

	sortedAr := make([]int, 0, len(ar))
	left, right := mergeSort(ar[:middle]), mergeSort(ar[middle:])

	var i, j = 0, 0
	for i < len(left) && j < len(right) {
		if left[i] > right[j] {
			sortedAr = append(sortedAr, right[j])
			j++
		} else {
			sortedAr = append(sortedAr, left[i])
			i++
		}
	}

	sortedAr = append(sortedAr, left[i:]...)
	sortedAr = append(sortedAr, right[j:]...)

	return sortedAr
}

// быстрая сортировка
func quickSort(ar []int) {
	if len(ar) < 2 {
		return
	}

	left, right := 0, len(ar) - 1
	pivotIndex := rand.Int() % len(ar)

	ar[pivotIndex], ar[right] = ar[right], ar[pivotIndex]

	for i := 0; i < len(ar); i++ {
		if ar[i] < ar[right] {
			ar[i], ar[left] = ar[left], ar[i]
			left++
		}
	}

	ar[left], ar[right] = ar[right], ar[left]

	quickSort(ar[:left])
	quickSort(ar[left + 1:])

	return
}


//сортировка вставкой
func insertionSort(ar []int) {
	if len(ar) < 2 {
		return
   }

	for i := 1; i < len(ar); i++ {
		for j := i; j > 0 && ar[j-1] > ar[j]; j-- {
			ar[j - 1], ar[j] = ar[j], ar[j-1]
	   }
	}
}
