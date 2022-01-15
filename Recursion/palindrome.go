package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	n := 1342
	ans := reverse(n)
	fmt.Println(ans)
	if n != ans {
		fmt.Println("Not Palindrome")
	} else {
		fmt.Println(" Palindrome")
	}
}

func reverse(n int) int {
	ans := 0
	var x int
	var rem int
	var length float64
	if n < 10 {
		return n
	}

	rem = n % 10
	fmt.Println(n, " ", rem)
	x = n / 10
	length = float64(len(strconv.Itoa(x)))
	fmt.Println(length)
	ans = (rem * int(math.Pow(10, length))) + reverse(x)
	fmt.Println(ans)

	return ans

}
