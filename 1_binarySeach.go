//WE KNOw ARRAY IS SORTED IN ASC ORDER
package main

import (
	"fmt"
)

func main() {
	f := fmt.Println
	// c := fmt.Print
	p := fmt.Printf

	// Entering length of array
	f("Enter Array Length")
	var n int
	fmt.Scanln(&n)

	// Input Array by user
	f("Enter Array (Should be ascending or descending)")
	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	f("Your array is: ", arr)

	// Entering length of array
	f("Enter Number to be searched in array")
	var numberToBeSearched int
	fmt.Scanln(&numberToBeSearched)

	// If ascending array then different function
	if arr[0] < arr[len(arr)-1] {
		ans := binarySearchAsc(arr, numberToBeSearched)
		if ans < 0 {
			p("%d is not present in entered array", numberToBeSearched)
			f()
		} else {
			p("%d is at index %d ", numberToBeSearched, ans)
			f()
		}

	} else {
		ans := binarySearchDesc(arr, numberToBeSearched)
		if ans < 0 {
			p("%d is not present in entered array", numberToBeSearched)
			f()
		} else {
			p("%d is at index %d ", numberToBeSearched, ans)
			f()
		}
	}

}

// If array is ascending
func binarySearchAsc(arr []int, target int) int {
	start := 0
	end := len(arr) - 1
	for start <= end {
		//	mid:=(start+end)/2   Might be possible k ye exceeed kr jaye int ki limit
		mid := start + (end-start)/2

		if arr[mid] > target {
			end = mid - 1
		} else if arr[mid] < target {
			start = mid + 1
		} else {
			return mid
		}

	}
	return -1
}

// If array is descending
func binarySearchDesc(arr []int, target int) int {
	start := 0
	end := len(arr) - 1
	for start <= end {
		//	mid:=(start+end)/2   Might be possible k ye exceeed kr jaye int ki limit
		mid := start + (end-start)/2

		if arr[mid] > target {
			start = mid + 1
		} else if arr[mid] < target {
			end = mid - 1
		} else {
			return mid
		}

	}
	return -1
}
