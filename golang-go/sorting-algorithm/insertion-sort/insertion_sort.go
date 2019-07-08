package main

import (
	"fmt"
	"math/rand"
	"time"
)

func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0; j-- {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			} else {
				break
			}
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
	result := InsertionSort(slice)

	//output after sort
	fmt.Printf("After:  %v\n", result)
}
