package main

import "fmt"

func main() {
	count(565000210)
	fmt.Println(countZeroes(565000200, 0))
}
func countZeroes(n int, c int) int {
	count := c
	rem := 0

	if n == 0 {
		return count
	}
	rem = n % 10
	if rem == 0 {
		return countZeroes(n/10, count+1)
	}

	return countZeroes(n/10, count)

}

func count(n int) {
	var rem, count int
	for n > 0 {
		rem = n % 10
		n /= 10
		if rem == 0 {
			count++
		}
	}
	fmt.Println(count)
}
