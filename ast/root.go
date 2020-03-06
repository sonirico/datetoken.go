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

func (rn *RootNode) String() string {
	var buf bytes.Buffer
	for _, node := range rn.Nodes {
		buf.WriteString(node.String())
	}
	return buf.String()
}

func (rn *RootNode) Inspect() string {
	return rn.String()
}
