package ast

import (
	"fmt"

	"github.com/sonirico/datetoken.go/token"
)

type NodeType string

const (
	Root       NodeType = "root"
	Value      NodeType = "value"
	Arithmetic NodeType = "arithmetic"
	Snap       NodeType = "snap"
)

type Node interface {
	Type() NodeType
	Literal() string
	Inspect() string
}

type ValueNode struct {
	Token token.Token
}

func (vn ValueNode) Type() NodeType {
	return Value
}

func (vn ValueNode) Literal() string {
	return vn.Token.Literal
}

func (vn ValueNode) Inspect() string {
	return vn.Literal()
}

type ArithmeticNode struct {
	Token token.Token

	Amount int64
	Sign   string
	Unit   string
}

func (an *ArithmeticNode) Type() NodeType {
	return Arithmetic
}

func (an *ArithmeticNode) Literal() string {
	return an.Token.Literal
}

func (an *ArithmeticNode) Inspect() string {
	return fmt.Sprintf("%s%d%s", an.Sign, an.Amount, an.Unit)
}

type SnapNode struct {
	Token token.Token

	Operator string

	Unit string
}

func (sn SnapNode) Type() NodeType {
	return Snap
}

func (sn SnapNode) Literal() string {
	return sn.Token.Literal
}

func (sn SnapNode) Inspect() string {
	return fmt.Sprintf("%s%s", sn.Operator, sn.Unit)
}
