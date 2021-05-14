package main

// Transform maze into graph.
// Each position represents a node
// You can walk to adjacent nodes with weight 1

// 011A1
// 11001
// 11011
// 01111
// B1001

// Returns graph from maze.
func MazeToGraph(maze [][]bool) *graph {
	cols := len(maze)
	rows := len(maze[0])

	g := Graph()

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if maze[i][j] {
				// N = uint(i*cols + j
				g.AddNode()
			}
		}
	}

	var n uint
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if maze[i][j] {

				if i != 0 {
					if !maze[i-1][j] {
						n = uint(i*cols + j)

						g.addEdge
					}
				}

				// N = uint(i*cols + j
				g.AddNode()
			}
		}
	}

	for n := range g.nodes {
		number = n.number
		i = n.number / uint(cols)
		j = n.number % uint(cols)
	}

	return g
}
