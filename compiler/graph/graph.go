package graph

import "github.com/matthewmueller/golly/compiler/decl"

// Graph struct
type Graph struct {
	nodes map[string]*Node
	edges map[string]string
}

// New function
func New() *Graph {
	return &Graph{
		nodes: map[string]*Node{},
		edges: map[string]string{},
	}
}

// AddDependency fn
func (g *Graph) AddDependency(parent decl.Declaration, child ...decl.Declaration) {
}

// Roots gets a list of root declarations (those without any dependants)
func (g *Graph) Roots() (nodes []Node) {
	return nodes
}

// Node struct
type Node struct {
	declaration decl.Declaration
}

// ID fn
func (n *Node) ID() string {
	return n.declaration.ID()
}

// Declaration fn
func (n *Node) Declaration() decl.Declaration {
	return n.declaration
}

// Dependencies fn
func (n *Node) Dependencies() (deps []Node) {
	return deps
}

// Sort topologically
func (n *Node) Sort() (list []Node) {
	return list
}