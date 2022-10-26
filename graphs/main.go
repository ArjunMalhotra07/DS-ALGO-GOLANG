package graphs

import "fmt"

func DetialsHelper() {
	f := fmt.Println
	f("Hello")
	details()

}

func details() {
	f := fmt.Println
	// p := fmt.Print
	s := fmt.Scanln
	var n, m int

	f("Enter number of Nodes and edges")
	s(&n, &m)

	matrixDetails(m)

}

func matrixDetails(length int) {
	f := fmt.Println
	p := fmt.Print
	s := fmt.Scanln

	var mat [6][6]int
	var e1, e2 int

	adjList := [6][]int{}
	f("Enter 6 Edges")

	for i := 0; i < length; i++ {
		s(&e1, &e2)
		mat[e1][e2] = 1
		mat[e2][e1] = 1
		adjList[e1] = append(adjList[e1], e2)

		adjList[e2] = append(adjList[e2], e1)
	}
	f()
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			p(mat[i][j])
			p(" ")
		}
		f()
	}
	f()
	f(adjList)
}

/*
1 2
1 3
2 4
3 4
3 5
4 5
*/
