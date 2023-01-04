package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const size = 50000
const threshold = 5000

func InsertSort(data []int) {
	i := 1
	for i < len(data) {
		h := data[i]
		j := i - 1
		for j >= 0 && h < data[j] {
			data[j+1] = data[j]
			j -= 1
		}
		data[j+1] = h
		i += 1
	}
}

func Partition(data []int) int {
	data[len(data)/2], data[0] = data[0], data[len(data)/2]
	pivot := data[0]
	mid := 0
	i := 1
	for i < len(data) {
		if data[i] < pivot {
			mid += 1
			data[i], data[mid] = data[mid], data[i]
		}
		i += 1
	}
	data[0], data[mid] = data[mid], data[0]
	return mid
}

func IsSorted(data []int) bool {
	for i := 1; i < len(data); i++ {
		if data[i] < data[i-1] {
			return false
		}
	}
	return true
}

func ConcurrentQuicksort(data []int, wg *sync.WaitGroup) {
	for len(data) >= 30 {
		mid := Partition(data)
		var portion []int
		if mid < len(data)/2 {
			portion = data[:mid]
			data = data[mid+1:]
		} else {
			portion = data[mid+1:]
			data = data[:mid]
		}
		if len(portion) > threshold {
			wg.Add(1)
			go func(data []int) {
				defer wg.Done()
				ConcurrentQuicksort(data, wg)
			}(portion)
		} else {
			ConcurrentQuicksort(portion, wg)
		}
	}
	InsertSort(data)
}

func QSort(data []int) {
	var wg sync.WaitGroup
	ConcurrentQuicksort(data, &wg)
	wg.Wait()
}

func partition(data []int, low, high int) int {
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
	data := make([]int, size)
	for i := 0; i < size; i++ {
		// data[i] = int(100 * rand.Float64())
		data[i] = rand.Intn(1000)
	}
	start := time.Now()
	QSort(data)
	elapsed := time.Since(start)
	fmt.Println("Elapsed time for concurrent quicksort = ", elapsed)
	fmt.Println("Is sorted: ", IsSorted(data))
}
