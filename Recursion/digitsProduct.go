package main

import "fmt"

func main() {
	n := 155
	if n < 0 {
		n *= -1
	}
	fmt.Print(product(n))
}

func product(n int) int {
	ans := 1
	if n == 0 {
		return ans
	}
	ans = n % 10
	n /= 10

	return ans * product(n)
}
