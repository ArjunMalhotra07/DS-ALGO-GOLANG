package main

import "fmt"

func main() {
	arr := []int{5, 6, 7, 8, 9, 10}
	fmt.Println(binarySearch(arr, 11, 0, len(arr)-1))

}

func binarySearch(arr []int, target int, start int, end int) int {

	if start > end {
		return -1
	}
	mid := start + (end-start)/2

	if target == arr[mid] {
		return mid
	}
	if arr[mid] < target {
		return binarySearch(arr, target, mid+1, end)
	}
	return binarySearch(arr, target, start, mid-1)

}
