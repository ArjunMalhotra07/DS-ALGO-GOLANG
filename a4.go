//Minimum element in a sorted and rotated array
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
	fmt.Println("Sorted array in Ascending order")
	print(arr, n)

	fmt.Println("Press 1: Left Rotation or 2: Right Rotation")
	var d int
	fmt.Scanln(&d)

	if d == 1 {
		fmt.Println("Rotate array by?")
		var e int
		fmt.Scanln(&e)
		leftRotate(arr, e, n)
	} else {
		fmt.Println("Rotate array by?")
		var e int
		fmt.Scanln(&e)
		rightRotate(arr, e, n)
	}

	print(arr, n)
	checkMinimum(arr, n)
}
func bSort(arr []int, n int) {
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
}
func print(arr []int, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d \n", arr[i])
	}
}

func leftRotate(arr []int, d int, n int) {
	for i := 0; i < d; i++ {
		leftRotatebyOne(arr, n)
	}
}
func leftRotatebyOne(arr []int, n int) {
	var i, temp int

	temp = arr[0]
	for i = 0; i < n-1; i++ {
		arr[i] = arr[i+1]
	}
	arr[n-1] = temp
}
func rightRotate(arr []int, d int, n int) {
	for i := 0; i < d; i++ {
		rightRotatebyOne(arr, n)
	}
}
func rightRotatebyOne(arr []int, n int) {
	var i, temp int

	temp = arr[n-1]
	for i = n - 1; i > 0; i-- {
		arr[i] = arr[i-1]

	}
	arr[0] = temp
}
func checkMinimum(arr []int, n int) {
	var temp int
	for i := 0; i < n-1; i++ {
		if arr[i] < arr[i+1] {
			temp = arr[i]
		}
	}
	fmt.Printf("smallest element is at %dth index", temp)

}
