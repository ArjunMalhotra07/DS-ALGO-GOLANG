package main

import "fmt"

func main() {
	var arr = []int{4, 8, 9, 101, 0, 1}
	SelectionSort(arr)
	fmt.Println(arr)
}

func SelectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		last := len(arr) - i - 1
		maxIndex := getMaxIndex(arr, 0, last)
		swap(arr, maxIndex, last)
	}
}

func swap(arr []int, maxIndex int, last int) {
	temp := arr[maxIndex]

	arr[maxIndex] = arr[last]
	arr[last] = temp

}

func getMaxIndex(arr []int, start int, last int) int {
	max := start
	end := last
	for i := start; i <= end; i++ {
		if arr[max] < arr[i] {
			max = i
		}
	}
	return max
}
