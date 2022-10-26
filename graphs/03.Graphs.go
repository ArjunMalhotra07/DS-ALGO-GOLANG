package graphs

import (
	"fmt"
)

//Graph structure
type Graph struct {
	vertices []*Vertex
}

//Vertex structure
type Vertex struct {
	key      int
	adjacent []*Vertex
}

//Contains
func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

//AddVertex
func (g *Graph) AddVertex(k int) {
	if contains(g.vertices, k) {
		err := fmt.Errorf("Vertex %v  already present", k)
		fmt.Println(err.Error())
	} else {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}

}

//AddEdge
func (g *Graph) AddEdge(from int, to int) {
	//getvertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)

	//check error
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("Error Adding Edge %v --> %v", from, to)
		fmt.Println(err.Error())
	} else if contains(fromVertex.adjacent, to) {
		err := fmt.Errorf("Existing Edge %v --> %v already exists", from, to)
		fmt.Println(err.Error())
	} else {
		//add edge
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}

//get vertex
func (g *Graph) getVertex(k int) *Vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

func GraphsHelper() {
	//f := fmt.Println
	test := &Graph{}

	test.AddVertex(100)
	test.AddVertex(200)
	test.AddVertex(300)
	test.AddVertex(400)
	test.AddVertex(500)

	test.AddEdge(200, 600)
	test.AddEdge(200, 300)
	test.AddEdge(200, 300)
	test.Print()
}
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\n Vertex %v : ", v.key)
		for _, v := range v.adjacent {
			fmt.Printf(" %v ", v.key)
		}
		fmt.Println()
	}
}
