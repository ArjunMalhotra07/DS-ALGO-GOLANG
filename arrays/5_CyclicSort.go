package arrays

import "fmt"

func SelectionSortHelper() {
	var arr = []int{2, 4, 5, 1, 3}
	Selectionsort(arr)
	fmt.Println(arr)
}

func Selectionsort(arr []int) {
	i := 0
	for i < len(arr) {
		correct := arr[i] - 1
		if arr[i] != arr[correct] {

			swap(arr, i, correct)

		} else {
			i++
		}
	}

}
func swap(arr []int, i int, correct int) {

	temp := arr[correct]
	arr[correct] = arr[i]
	arr[i] = temp
}
