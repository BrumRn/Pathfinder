package main

import "fmt"

func main() {

	g := Graph()
	n1 := g.AddNode(1)
	n2 := g.AddNode(2)
	n3 := g.AddNode(3)
	n33 := g.AddNode(4)
	n4 := g.AddNode(5)
	n5 := g.AddNode(6)

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

	maze := [][]bool{
		{true, true, true, true},
		{true, false, false, true},
		{true, false, true, true},
		{true, true, true, false},
	}

	g = MazeToGraph(maze)
	sol = solution{queue: g}

	//n := (2)*len(maze) + (3)

	sol.Dijkstras(g.nodes[0])

	fmt.Println(sol.dist)

}
