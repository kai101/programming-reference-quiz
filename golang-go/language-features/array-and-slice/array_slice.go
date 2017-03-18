package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Array and Slice are 2 important concept in GO!\n")

	var arrayA [5]int = [5]int{1, 2, 3, 4, 5} //this is literal array.
	var sliceA []int                          // this is slice

	fmt.Printf("1. Type of Array is %T while type of Slice is %T\n", arrayA, sliceA)

	fmt.Printf("While they seems similar, array is an actual list of element, while slice is a view to the array.\n")

	//lets make "slice" become a view for "array"
	sliceA = arrayA[0:4] //slice will now have a full view of the arrayA.

	var sliceB []int     //
	sliceB = arrayA[1:4] //sliceB will now have view of index 1 to 4 of the arrayA, index 0 is excluded.

	fmt.Printf("2. Lets check the length len() of sliceA: %v and sliceB: %v\n", len(sliceA), len(sliceB))
	fmt.Printf("Beside length, Slice have another properties call Capacity cap().\n")
	fmt.Printf("3. Capacity of sliceA: %v and sliceB: %v\n", cap(sliceA), cap(sliceB))

	//changing array data
	fmt.Printf("As mentioned before, Slice is a view of Array, any changes on the Array will reflect on Slice.\n")

	arrayA[2] = 1000 //changing the arrayA will show on sliceA and sliceB

	//sliceB is started at index 1 of arrayA arrayA[2] equal to sliceB[1]
	fmt.Printf("4. Lets check the value of sliceA[2]: %v and sliceB[1]: %v\n", sliceA[2], sliceB[1])

	//changing slice data.
	fmt.Printf("Vice versa, changing a slice will access the array and change the corresponding data.\n")

	sliceA[2] = 3000

	fmt.Printf("5. Lets check the value of arrayA[2]: %v and sliceB[1]: %v\n", arrayA[2], sliceB[1])
	
	//TODO fix size of array, dynamic size of slice etc.

}
