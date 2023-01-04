package main

import (
	"fmt"
)

type Ordered interface {
	~float64 | ~int | ~string
}

func quicksort[T Ordered](data []T, low, high int) {
	if low < high {
		var pivot = partition(data, low, high)
		quicksort(data, low, pivot)
		quicksort(data, pivot+1, high)
	}
}

func partition[T Ordered](data []T, low, high int) int {
	//Pick a lowest bound element as a pivot value
	var pivot = data[low]
	var i = low
	var j = high

	for i < j {
		for data[i] <= pivot && i < high {
			i++
		}

		for data[j] > pivot && j > low {
			j--
		}

		if i < j {
			data[i], data[j] = data[j], data[i]
		}
	}

	data[low] = data[j]
	data[j] = pivot
	return j
}

func main() {
	numbers := []float64{3.5, -2.4, 12.8, 9.1}
	names := []string{"Zachary", "John", "Moe", "Jim", "Robert"}
	quicksort[float64](numbers, 0, len(numbers)-1)
	fmt.Println(numbers)
	quicksort[string](names, 0, len(names)-1)
	fmt.Println(names)
}
