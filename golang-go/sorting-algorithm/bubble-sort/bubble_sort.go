package main

import (
	"fmt"
	"math/rand"
	"time"
)

func bubbleSort(arr []int) []int{
	for i := len(arr) - 1 ; i > 0 ; i--{
		for j := 0; j < i; j++   {
			if(arr[j] > arr[j+1]){
				temp := arr[j+1];
				arr[j+1] = arr[j];
				arr[j] = temp;
			}
		}
	}
	return arr;
}

func main() {
	var arr [10]int
	var slice = arr[:] //make a slice #see golang slice and array
	
	//setup pseudorandom number generator
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	
	//initialize array with random number of 0 to 100;
	for index := range arr{
		arr[index] = random.Intn(100);
	}
	
	//output before sort
	fmt.Printf("Before: %v\n", arr)
	
	//sorting
	bubbleSort(slice);
	
	//output after sort
	fmt.Printf("after:  %v\n", arr)
}