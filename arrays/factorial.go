package main

import "fmt"

func main() {
	fmt.Println(fact(15))
	fmt.Println(fact(5))
}
func fact(n int) int {
	if n <= 1 {
		return 1
	}
	return n * fact(n-1)
}

func fact1(n int) int {
	p := 1
	for i := 1; i <= n; i++ {
		p *= i
	}
	return p
}
