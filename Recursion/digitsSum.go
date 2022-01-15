package main

import "fmt"

func main() {
	ans := fact(1342)
	fmt.Println(ans)
}

func fact(n int) int {
	ans := 0
	var q int
	if n <= 1 {
		return 1
	}
	ans = n % 10
	q = n / 10

	return ans + fact(q)
}
