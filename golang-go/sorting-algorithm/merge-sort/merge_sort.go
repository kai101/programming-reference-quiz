package main

import (
	"fmt"
	"math/rand"
	"time"
)

func MergeSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	var mid int = len(arr) / 2
	var left, right []int
	
	//make seperate copy of left and right for manipulation on arr.
	left = make([]int, len(arr[:mid]))
	copy(left, MergeSort(arr[:mid])) //recursive call
	
	right = make([]int, len(arr[mid:]))
	copy(right, MergeSort(arr[mid:])) //recursive call

	var i, j, k int = 0, 0, 0

	//smart merging by having 2 pointers on left and right and compare them.
	for (i < len(left)) && (j < len(right)) {
		if left[i] < right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}

	//process the remaining elements in left
	for i < len(left) {
		arr[k] = left[i]
		i++
		k++
	}

	//process the remaining elements in right
	for j < len(right) {
		arr[k] = right[j]
		j++
		k++
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
	result := MergeSort(slice)

	//output after sort
	fmt.Printf("After:  %v\n", result)
}
