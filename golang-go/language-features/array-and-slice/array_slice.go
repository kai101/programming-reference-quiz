package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Array and Slice are 2 important concept in GO!\n")

	var arrayA [5]int = [5]int{1, 2, 3, 4, 5} //this is literal array.
	var sliceA []int                          // this is slice

	fmt.Printf("1. Type of Array is %T while type of Slice is %T\n", arrayA, sliceA)
	//1. Type of Array is [5]int while type of Slice is []int
	fmt.Printf("While they seems similar, array is an actual list of element, while slice is a view to the array.\n")
	//While they seems similar, array is an actual list of element, while slice is a view to the array.

	//lets make "slice" become a view for "array"
	sliceA = arrayA[0:4] //slice will now have a full view of the arrayA.

	var sliceB []int     //
	sliceB = arrayA[1:4] //sliceB will now have view of index 1 to 4 of the arrayA, index 0 is excluded.

	fmt.Printf("2. Lets check the length len() of sliceA: %v and sliceB: %v\n", len(sliceA), len(sliceB))
	//2. Lets check the length len() of sliceA: 4 and sliceB: 3
	fmt.Printf("Beside length, Slice have another properties call Capacity cap().\n")
	//Beside length, Slice have another properties call Capacity cap().
	fmt.Printf("3. Capacity of sliceA: %v and sliceB: %v\n", cap(sliceA), cap(sliceB))
	//3. Capacity of sliceA: 5 and sliceB: 4

	//changing array data
	fmt.Printf("As mentioned before, Slice is a view of Array, any changes on the Array will reflect on Slice.\n")
	//As mentioned before, Slice is a view of Array, any changes on the Array will reflect on Slice.

	arrayA[2] = 1000 //changing the arrayA will show on sliceA and sliceB

	//sliceB is started at index 1 of arrayAi, therefore  arrayA[2] equal to sliceB[1]
	fmt.Printf("4. Lets check the value of sliceA[2]: %v and sliceB[1]: %v\n", sliceA[2], sliceB[1])
	//4. Lets check the value of sliceA[2]: 1000 and sliceB[1]: 1000

	//changing slice data.
	fmt.Printf("Vice versa, changing a slice will access the array and change the corresponding data.\n")

	sliceA[2] = 3000

	fmt.Printf("5. Lets check the value of arrayA[2]: %v and sliceB[1]: %v\n", arrayA[2], sliceB[1])
	//5. Lets check the value of arrayA[2]: 3000 and sliceB[1]: 3000

	//array size are fix, but slice cap
	//You cant do this
	//arrayA = append(arrayA, 9)

	//but you can append slice, remember slice is a reference to array.
	sliceA = append(sliceA, 9)
	fmt.Printf("len=%d cap=%d %v\n", len(sliceA), cap(sliceA), sliceA)

	//however sliceB is not affected.
	fmt.Printf("len=%d cap=%d %v\n", len(sliceB), cap(sliceB), sliceB)

	//sliceA cap is 5 and len is 5, what if we append one more number to sliceA
	sliceA = append(sliceA, 10)
	fmt.Printf("len=%d cap=%d %v\n", len(sliceA), cap(sliceA), sliceA)

	fmt.Printf("There you go, slice is a reference, and dynamic component for a Array. And do not forget array is fixed length.")
}
