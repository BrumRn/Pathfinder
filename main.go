package main

import "fmt"

func main() {
	maze := [][]bool{
		{true, true, true, true},
		{true, false, false, true},
		{true, false, true, true},
		{true, true, true, false},
	}

	dist, prev := FindShortestPath(maze, 0, 0, 2, 3)

	fmt.Println(dist, prev)

}
