package ast

type NodeType string

const (
	Root   NodeType = "root"
	Amount NodeType = "amount"
	Unit   NodeType = "unit"
)

type Node interface {
	Type() NodeType
	String() string
	Inspect() string
}
