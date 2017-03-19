package main

import (
	"fmt"
	"math/rand"
	"time"
)

func BucketSort(arr []int) []int {
	var largest int
	var bucket, result []int

	//initialize result
	resultLens := len(arr)
	result = make([]int, resultLens, resultLens)

	//determine largest integer
	for _, v := range arr {
		if v > largest {
			largest = v
		}
	}

	//initialize bucket
	bucket = make([]int, largest+1, largest+1)

	//add value into bucket
	for _, v := range arr {
		bucket[v]++
	}

	//count the number inside bucket and output to result.
	resultPointer := 0
	for i, v := range bucket {
		for j := 0; j < v; j++ {
			result[resultPointer] = i
			resultPointer++
		}
	}

	return result
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
	result := BucketSort(slice)

	//output after sort
	fmt.Printf("After:  %v\n", result)
}
