package main

import (
	"fmt"
	"math/rand"
	"time"
)

func swap(arr []int, i,j int){
	var temp = arr[i];
	arr[i] = arr[j];
	arr[j] = temp;
}

func quickSort(arr []int, left, right int) []int{
	var pivot, parIndex int
	
	if(left < right){
		pivot = right
		parIndex = left
		
		for i := left;i<right;i++{
			if(arr[i] < arr[pivot]){
				swap(arr,i,parIndex)
				parIndex++
			}
		}
		
		swap(arr, parIndex,pivot)
		
		quickSort(arr, left, parIndex -1)
		quickSort(arr, parIndex +1 , right)
	
	}
	return arr;
}

func QuickSort(arr []int) []int {
	return quickSort(arr,0,len(arr)-1)
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
	result := QuickSort(slice)

	//output after sort
	fmt.Printf("after:  %v\n", result)
}
