package models

import (
	"bytes"
	"errors"
	"fmt"
)

var EmptyTokenError = errors.New("empty tokens default to 'just now' value. You can still read value to my left")

type InvalidTokenError struct {
	Literal string
	Errors  []string
}

func NewInvalidTokenError(literal string, errors []string) InvalidTokenError {
	return InvalidTokenError{Literal: literal, Errors: errors}
}

func (ite InvalidTokenError) Error() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("invalid token '%s'", ite.Literal))
	if len(ite.Errors) > 0 {
		for i, e := range ite.Errors {
			buf.WriteString(fmt.Sprintf("\n %d. %s", i, e))
		}
	}
	return buf.String()
}
