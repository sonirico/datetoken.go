package ast

type NodeType string

const (
	Root       NodeType = "root"
	Value      NodeType = "value"
	Arithmetic NodeType = "arithmetic"
	Snap       NodeType = "snap"
)

// Node interface defines the structure of all ast nodes
type Node interface {
	// Type returns the ast node type
	Type() NodeType
	// Literal returns the ast node main token value
	Literal() string
	// Inspect returns the ast node string presentation
	Inspect() string
}
