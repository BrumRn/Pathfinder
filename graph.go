package main

// Implement a graph.
// Olny allows nonnegative values
type graph struct {
	nodes map[uint]*node
	size  uint
}

// Store node with adjacency list.
type node struct {
	number uint
	name   uint
	edges  []*edge
}

// Store edge with target and weight.
type edge struct {
	target *node
	weight uint
}

// Initializes an empty graph.
func Graph() *graph {
	n := make(map[uint]*node)
	return &graph{nodes: n}
}

// Adds new node to graph
func (g *graph) AddNode(name uint) *node {
	n := node{number: g.size, name: name}
	g.nodes[name] = &n
	g.size++
	return &n
}

// Adds edge to node.
func (n *node) AddEdge(target *node, weight uint) {
	e := edge{target: target, weight: weight}
	n.edges = append(n.edges, &e)
}
