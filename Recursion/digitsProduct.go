package main

import "fmt"

func main() {
	fmt.Print(reverse(1234))
}
func reverse(n int) int {
	var ans int = 0
	var r int = 0

	ans = n / 10
	r = n % 10

	return (n%10)*10 + reverse(ans)
}
