package main

import "fmt"

func main() {
	print(1)
}

func print(n int) {

	if n == 5 {

		return
	}
	fmt.Println(n)

	print(n + 1)
}
