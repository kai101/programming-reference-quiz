package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Heapify function
func Heapify(arr []int, size int, index int) {
	left := index*2 + 1
	right := index*2 + 2
	largest := index

	if left < size && arr[left] > arr[largest] {
		largest = left
	}

	if right < size && arr[right] > arr[largest] {
		largest = right
	}

	if largest != index {
		arr[index], arr[largest] = arr[largest], arr[index]
		Heapify(arr, size, largest)
	}
}

// HeapSort function
func HeapSort(arr []int) []int {
	var size = len(arr)

	// creating heap
	for i := size/2 - 1; i >= 0; i-- {
		//
		Heapify(arr, size, i)
	}

	// move root to the end and heapify again.
	for i := size - 1; i >= 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		Heapify(arr, i, 0)
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
	result := HeapSort(slice)

	//output after sort
	fmt.Printf("After:  %v\n", result)
}
