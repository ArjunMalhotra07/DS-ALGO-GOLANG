package main

import "fmt"

func main() {
	var arr = []int{5, 4, 3, 2, 1}
	BubbleSort(arr)
	fmt.Println(arr)
}

func BubbleSort(arr []int) {
	var b bool
	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(arr)-i; j++ {
			if arr[j] < arr[j-1] {
				temp := arr[j-1]
				arr[j-1] = arr[j]
				arr[j] = temp
				b = true
			}
		}
	}
	if b == false {
		return
	}
}
