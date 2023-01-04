package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Comparing two isSorted function

const size = 100_000_000

var data []float64

func isSorted1(data []float64) bool {
	var data1 []float64
	data1 = make([]float64, len(data))

	copy(data1, data) // Copies data into data1
	sort.Float64s(data1)

	//Compare data and data1
	for i := 0; i < size; i++ {
		if data[i] != data1[i] {
			return false
		}
	}

	return true
}

func isSorted2(data []float64) bool {
	for i := 1; i < len(data); i++ {
		if data[i] < data[i-1] {
			return false
		}
	}
	return true
}

func main() {
	data = make([]float64, size)
	for i := 0; i < size; i++ {
		data[i] = 100.0 * rand.Float64()
	}

	start := time.Now()
	result := isSorted1(data)
	elapsed := time.Since(start)
	fmt.Println("Sorted: ", result)
	fmt.Println("elapsed using sorted1:", elapsed)

	data2 := make([]float64, size)
	for i := 0; i < size; i++ {
		data2[i] = float64(2 * i)
	}

	start = time.Now()
	result = isSorted1(data2)
	elapsed = time.Since(start)
	fmt.Println("Sorted: ", result)
	fmt.Println("elapsed using sorted1:", elapsed)

	start = time.Now()
	result = isSorted2(data)
	elapsed = time.Since(start)
	fmt.Println("\nSorted: ", result)
	fmt.Println("elapsed using sorted2", elapsed)

	start = time.Now()
	result = isSorted2(data2)
	elapsed = time.Since(start)
	fmt.Println("Sorted: ", result)
	fmt.Println("elapsed using sorted2:", elapsed)
}
