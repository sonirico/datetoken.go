package ast

type NodeType string

const (
	Amount NodeType = "amount"
	Unit   NodeType = "unit"
)

type Node interface {
	Type() NodeType
	String() string
	Inspect() string
}
