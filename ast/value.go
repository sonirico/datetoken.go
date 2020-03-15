package ast

import (
	"github.com/sonirico/datetoken.go/token"
)

// ValueNode represents literals associated with values, such as "now" or "yesterday"
type ValueNode struct {
	Token token.Token
}

// Type returns the ast node type
func (vn ValueNode) Type() NodeType {
	return Value
}

// Literal returns the ast node main token value
func (vn ValueNode) Literal() string {
	return vn.Token.Literal
}

// Inspect returns the ast node string presentation
func (vn ValueNode) Inspect() string {
	return vn.Literal()
}
