//Bubble Sort
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

	bSort(arr, n)
	fmt.Println("Sorted array in Descending order")
	print(arr, n)

}
func bSort(arr []int, n int) {
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if arr[j] < arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
}
func print(arr []int, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}
}
