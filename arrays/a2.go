//Linear Search
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

	fmt.Println("Enter the element you want to search ")
	var e int
	fmt.Scanln(&e)
	linearSearch(arr, n, e)

}
func linearSearch(arr []int, n int, e int) {
	for i := 0; i < n; i++ {
		if arr[i] == e {
			fmt.Printf("%d found at %dth position", e, i)
		}
	}
	fmt.Println("Element not found")
}
