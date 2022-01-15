package main

import "fmt"

func main() {
	first := 0
	second := 1
	n := 100
	fibonacci(first, second, n)
	fmt.Println()
	fibonacciTerms(first, second, 15)
	fmt.Println()
	ans := fibonacciNumber(10)
	fmt.Println(ans)
}

//Uptil N fibonacci series
func fibonacci(first int, second int, end int) {
	var z int = 0
	x := first
	y := second
	z = x + y
	n := end

	for z > n {
		return
	}

	fmt.Printf("%v ", z)

	x = y
	y = z
	fibonacci(x, y, n)

}

//N number of Fibonacci Series
func fibonacciTerms(first int, second int, end int) {
	var z int = 0
	x := first
	y := second
	z = x + y

	for end == 0 {
		return
	}

	fmt.Printf("%v ", z)

	x = y
	y = z
	fibonacciTerms(x, y, end-1)

}

//Nth Fibonacci Number
func fibonacciNumber(n int) int {

	if n < 2 {
		return n
	}

	return fibonacciNumber(n-1) + fibonacciNumber(n-2)

}
