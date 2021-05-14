package main

import "fmt"

func main() {

	g := Graph()
	n1 := g.AddNode()
	n2 := g.AddNode()
	n3 := g.AddNode()
	n33 := g.AddNode()
	n4 := g.AddNode()
	n5 := g.AddNode()

	n1.AddEdge(n2, 1)
	n1.AddEdge(n3, 5)
	n2.AddEdge(n3, 1)
	n2.AddEdge(n4, 10)
	n3.AddEdge(n4, 8)
	n33.AddEdge(n1, 1)
	n4.AddEdge(n1, 1)
	n5.AddEdge(n1, 9)

	sol := solution{queue: g}

	sol.Dijkstras(n1)
	fmt.Println(sol.dist)
}
