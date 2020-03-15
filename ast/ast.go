package ast

// NodeType represents the kind of the ast node category
type NodeType string

const (
	// Root is the root node in the AST
	Root NodeType = "root"
	// Value is a value node in the AST, such as "now" or "yesterday"
	Value NodeType = "value"
	// Arithmetic is a arithmetic node in the AST, to add or subtract dates
	Arithmetic NodeType = "arithmetic"
	// Snap is a snap node in the AST, to move the clock to start or end of other dates
	Snap NodeType = "snap"
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
