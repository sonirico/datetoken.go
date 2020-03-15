package ast

import (
	"bytes"
)

// RootNode holds the rest of the ast tree
type RootNode struct {
	Nodes []Node
}

// NewRootNode returns a new instance of RootNode
func NewRootNode() *RootNode {
	return &RootNode{
		Nodes: []Node{},
	}
}

// Type returns the ast node type
func (rn *RootNode) Type() NodeType {
	return Root
}

func (rn *RootNode) addNode(node Node) {
	rn.Nodes = append(rn.Nodes, node)
}

// Literal returns the ast node main token value
func (rn *RootNode) Literal() string {
	return ""
}

// Inspect returns the ast node string presentation
func (rn *RootNode) Inspect() string {
	var buf bytes.Buffer
	for _, node := range rn.Nodes {
		buf.WriteString(node.Inspect())
	}
	return buf.String()
}
