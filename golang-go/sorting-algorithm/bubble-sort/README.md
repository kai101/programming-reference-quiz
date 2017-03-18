# Bubble Sort - golang/go

Simple bubble sort, probably your first ever sorting algorithm. 

## Specification
Worst run time O(n^2)

## Algorithm Explanation

Bubble sort uses an analogy of a big bubble rises from bottom of water to the surface of water. In terms of data, an element with larger value will move to the end of array.

## Procedure
1. Compare first element with its neigbhor, the 2nd element.
2. If the first element value is larger than 2nd element, swap their position.
3. Compare 2nd element with third element, if 2nd element is larger than third element, swap their position.
4. Continue until the last element is compared. 
5. At this point, the last element 'n' is the largest element.
6. Repeat step 1 until the n-1 element. After the 2nd loop, the 2nd largest element is determined. Repeat step 1 again until n-2 to determine the 3rd largest element. Continue until all elements are determined.