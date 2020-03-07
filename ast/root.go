package ast

import (
	"bytes"
)

type RootNode struct {
	Nodes []Node
}

func NewRootNode() *RootNode {
	return &RootNode{
		Nodes: []Node{},
	}
}

func (rm *RootNode) Type() NodeType {
	return Root
}

func (rn *RootNode) AddNode(node Node) {
	rn.Nodes = append(rn.Nodes, node)
}

func (rn *RootNode) Literal() string {
	return ""
}

func (rn *RootNode) Inspect() string {
	var buf bytes.Buffer
	for _, node := range rn.Nodes {
		buf.WriteString(node.Inspect())
	}
	return buf.String()
}
