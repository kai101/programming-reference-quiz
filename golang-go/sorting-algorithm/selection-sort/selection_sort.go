package main

import (
	"fmt"
	"math/rand"
	"time"
)

// SelectionSort function
func SelectionSort(arr []int) []int {
	min := 0
	for i := range arr {
		min = i

		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				min = j
			}
		}

		//swap
		if min != i {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}

	return arr
}

func main() {
	var arr [10]int
	var slice = arr[:] //make a slice #see golang slice and array

	//setup pseudorandom number generator
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	//initialize array with random number of 0 to 100;
	for index := range arr {
		arr[index] = random.Intn(100)
	}

	//output before sort
	fmt.Printf("Before: %v\n", arr)

	//sorting
	result := SelectionSort(slice)

	//output after sort
	fmt.Printf("After:  %v\n", result)
}
