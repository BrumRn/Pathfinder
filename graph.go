package main

// Implement a graph.
// Olny allows nonnegative values
type graph struct {
	nodes map[*node]bool
	size  uint
}

// Store node with adjacency list.
type node struct {
	number uint
	edges  []*edge
}

// Store edge with target and weight.
type edge struct {
	target *node
	weight uint
}

// Initializes an empty graph.
func Graph() *graph {
	n := make(map[*node]bool)
	return &graph{nodes: n}
}

// Adds new node to graph
func (g *graph) AddNode() *node {
	n := node{number: g.size}
	g.nodes[&n] = true
	g.size++
	return &n
}

// Adds edge to node.
func (n *node) AddEdge(target *node, weight uint) {
	e := edge{target: target, weight: weight}
	n.edges = append(n.edges, &e)
}
