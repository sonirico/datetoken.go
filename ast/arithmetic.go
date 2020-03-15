package ast

import (
	"fmt"
	"github.com/sonirico/datetoken.go/token"
)

// ArithmeticNode represents addition or subtraction operations on dates
type ArithmeticNode struct {
	Token token.Token

	Amount int64
	Sign   string
	Unit   string
}

// Type returns the ast node type
func (an *ArithmeticNode) Type() NodeType {
	return Arithmetic
}

// Literal returns the ast node main token value
func (an *ArithmeticNode) Literal() string {
	return an.Token.Literal
}

// Inspect returns the ast node string presentation
func (an *ArithmeticNode) Inspect() string {
	return fmt.Sprintf("%s%d%s", an.Sign, an.Amount, an.Unit)
}
