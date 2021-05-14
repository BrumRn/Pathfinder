package main

// Store graph of nodes.
type solution struct {
	queue *graph
	dist  []uint
	prev  []*node
	inf   uint
}

// Return shortest path from s to all other nodes.
// Algorithm returns inf if unconnected.
func (g *solution) Dijkstras(s *node) {
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
