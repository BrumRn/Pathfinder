package main

import "fmt"

// Store graph of nodes.
type solution struct {
	queue *graph
	dist  []uint
	prev  []*node
	inf   uint
}

// Finds shortest path.
func FindShortestPath(maze [][]uint, iS uint, jS uint, iT uint, jT uint) ([]uint, []*node) {
	queue := MazeToGraph(maze)
	g := solution{queue: queue}

	start := g.queue.nodes[iS*g.queue.cols+jS]
	target := g.queue.nodes[iT*g.queue.cols+jT]
	g.dijkstras(start, target)
	path := g.constructSolution(start, target)
	g.printPath(maze, path)
	return g.dist, g.prev
}

// Print path
func (g *solution) printPath(maze [][]uint, path []uint) {
	var i uint
	var j uint
	for _, n := range path {
		i = n / uint(len(maze))
		j = n % uint(len(maze))
		maze[i][j] = 2
	}

	maze[path[0]/uint(len(maze))][path[0]%uint(len(maze))] = 3
	maze[path[len(path)-1]/uint(len(maze))][path[len(path)-1]%uint(len(maze))] = 4

	for i := 0; i < len(maze[0]); i++ {
		for j := 0; j < len(maze); j++ {
			if maze[i][j] == 4 {
				fmt.Print("A" + " ")
			} else if maze[i][j] == 3 {
				fmt.Print("B" + " ")
			} else if maze[i][j] == 2 {
				fmt.Print("X" + " ")
			} else {
				fmt.Print(fmt.Sprint(maze[i][j]) + " ")
			}

		}
		fmt.Println()
	}

}

// Reconstructs best solution
func (g *solution) constructSolution(start *node, target *node) []uint {
	var path []uint
	n := target
	path = append(path, target.name)
	for n != start {
		n = g.prev[n.number]
		path = append(path, n.name)
	}
	return path
}

// Return shortest path from s to all other nodes.
// Algorithm returns inf if unconnected.
func (g *solution) dijkstras(s *node, target *node) {
	g.dist = make([]uint, len(g.queue.nodes))
	g.prev = make([]*node, len(g.queue.nodes))
	g.inf = ^uint(0)

	for _, v := range g.queue.nodes {
		g.dist[v.number] = g.inf
	}
	g.dist[s.number] = 0

	var u *node
	for len(g.queue.nodes) > 0 {
		u = g.closestNode()

		if u == nil {
			break
		}
		delete(g.queue.nodes, u.name)

		if u == target {
			break
		}

		var alt uint
		var v *node
		for _, e := range u.edges {
			alt = g.dist[u.number] + e.weight
			v = e.target
			if alt < g.dist[v.number] {
				g.dist[v.number] = alt
				g.prev[v.number] = u
			}
		}
	}
}

// Find closest node.
func (g *solution) closestNode() *node {
	minDist := g.inf
	var minNode *node
	for _, v := range g.queue.nodes {
		if g.dist[v.number] < minDist {
			minDist = g.dist[v.number]
			minNode = v
		}
	}
	return minNode
}
