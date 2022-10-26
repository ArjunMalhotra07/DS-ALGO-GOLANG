//Find smallest number greater than or equal to Target Integer
//Same can be done for greatest number smaller or equal to Target Integer

package arrays

import "fmt"

func AscHelper() {
	fmt.Println("Hi")

	var x int = 5
	var arr = make([]int, x)
	for i := 0; i < x; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	fmt.Println(arr)
	if arr[0] < arr[len(arr)-1] {
		ans := asc(arr, 0)
		fmt.Println("Ans: ", ans)
	} else {
		ans := desc(arr, 4)
		fmt.Println("Ans: ", ans)
	}
}

func asc(arr []int, target int) int {
	start := 0
	end := len(arr) - 1

	if target > arr[len(arr)-1] || target < arr[0] {
		return -1
	}

	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] > target {
			end = mid - 1
		} else if arr[mid] < target {
			start = mid + 1
		} else {
			return mid
		}
	}
	return arr[start]
}

func desc(arr []int, target int) int {
	start := 0
	end := len(arr) - 1
	if target < arr[len(arr)-1] || target > arr[0] {
		return -1
	}
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
	return arr[end]
}
