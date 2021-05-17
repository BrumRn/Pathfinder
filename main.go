package main

import (
	"fmt"
	"math/rand"
)

// Generates random maze
func randomMaze(n uint) [][]uint {
	maze := make([][]uint, n)
	for i := 0; i < int(n); i++ {
		maze[i] = make([]uint, n)
	}

	for i := 0; i < len(maze[0]); i++ {
		for j := 0; j < len(maze); j++ {
			k := uint(rand.Intn(4.))
			if k == 0 {
				maze[i][j] = 0
			} else {
				maze[i][j] = 1
			}
		}
	}

	return maze
}

func main() {
	maze := [][]uint{
		{1, 1, 1, 1},
		{1, 0, 0, 1},
		{1, 0, 1, 1},
		{1, 1, 1, 0},
	}

	var n uint = 25
	maze = randomMaze(n)
	for i := 0; i < len(maze[0]); i++ {
		for j := 0; j < len(maze); j++ {
			fmt.Print(fmt.Sprint(maze[i][j]) + " ")
		}
		fmt.Println()
	}
	fmt.Println()
	dist, prev := FindShortestPath(maze, 0, n/2, n-1, n/2)

	fmt.Println(dist, prev)

}
