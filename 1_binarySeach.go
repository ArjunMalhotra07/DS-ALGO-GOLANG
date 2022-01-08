//WE KNOE ARRAY IS SORTED IN ASC ORDER
package main

import "fmt"

func main() {

	var n int
	fmt.Println("Enter Array Length")
	fmt.Scanln(&n)
	fmt.Println("Enter Array")
	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	fmt.Println("Your array is: ", arr)

	if arr[0] < arr[len(arr)-1] {
		ans := binarySearchAsc(arr, 50)
		fmt.Println(ans)
	} else {
		ans := binarySearchDesc(arr, -1)
		fmt.Println(ans)
	}

}

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
