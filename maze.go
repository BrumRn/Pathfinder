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
	g.cols = uint(cols)
	g.rows = uint(rows)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if maze[i][j] {
				g.AddNode(uint(i*cols + j))
			}
		}
	}

	var n *node
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {

			if maze[i][j] {
				n = g.nodes[uint(i*cols+j)]

				if i != 0 {
					if maze[i-1][j] {
						n.AddEdge(g.nodes[uint((i-1)*cols+j)], 1)
					}
				}

				if i != rows-1 {
					if maze[i+1][j] {
						n.AddEdge(g.nodes[uint((i+1)*cols+j)], 1)
					}
				}

				if j != 0 {
					if maze[i][j-1] {
						n.AddEdge(g.nodes[uint(i*cols+j-1)], 1)
					}
				}

				if j != rows-1 {
					if maze[i][j+1] {
						n.AddEdge(g.nodes[uint(i*cols+j+1)], 1)
					}
				}
			}
		}
	}

	return g
}
