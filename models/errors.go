package models

import (
	"bytes"
	"errors"
	"fmt"
)

// ErrEmptyToken is yielded when the parser returns no nodes
var ErrEmptyToken = errors.New("empty tokens default to 'just now' value. You can still read value to my left")

// ErrInvalidToken is yielded when the parser encounters illegal tokens on the payload
type ErrInvalidToken struct {
	Literal string
	Errors  []string
}

// NewInvalidTokenError returns a new instance of ErrInvalidToken
func NewInvalidTokenError(literal string, errors []string) ErrInvalidToken {
	return ErrInvalidToken{Literal: literal, Errors: errors}
}

// Error returns the string representation of the error
func (ite ErrInvalidToken) Error() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("invalid token '%s'", ite.Literal))
	if len(ite.Errors) > 0 {
		for i, e := range ite.Errors {
			buf.WriteString(fmt.Sprintf("\n %d. %s", i, e))
		}
	}
	return buf.String()
}
