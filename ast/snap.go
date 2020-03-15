package ast

import (
	"fmt"
	"github.com/sonirico/datetoken.go/token"
)

// SnapNode represents the action of moving the clock to the start or the end of the time unit it holds
type SnapNode struct {
	Token token.Token

	Operator string

	Unit string
}

// Type returns the ast node type
func (sn SnapNode) Type() NodeType {
	return Snap
}

// Literal returns the ast node main token value
func (sn SnapNode) Literal() string {
	return sn.Token.Literal
}

// Inspect returns the ast node string presentation
func (sn SnapNode) Inspect() string {
	return fmt.Sprintf("%s%s", sn.Operator, sn.Unit)
}
