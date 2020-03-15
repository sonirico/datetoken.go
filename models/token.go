package models

import (
	"bytes"
	"fmt"
	"time"
)

// Token encapsulates the input and output of a token evaluation
type Token struct {
	literal string
	date    time.Time
	errors  []error
}

// NewToken returns a new instance of Token
func NewToken(payload string, date time.Time) Token {
	return Token{
		literal: payload,
		date:    date,
		errors:  make([]error, 0),
	}
}

// AddError adds errors to the error list of the token
func (t Token) AddError(err error) {
	t.errors = append(t.errors, err)
}

// Literal returns the token as given by the client
func (t Token) Literal() string {
	return t.literal
}

// String returns debug information of the token
func (t Token) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("payload: '%s'", t.literal))
	buf.WriteString("\n")
	buf.WriteString(fmt.Sprintf("date: '%s'", t.date.String()))
	buf.WriteString("\n")
	buf.WriteString(t.ErrorString())
	return buf.String()
}

// Errors will return the list of errors encountered during the evaluation
func (t Token) Errors() []error {
	return t.errors
}

// ErrorString will return a formatted string version of the errors
func (t Token) ErrorString() string {
	var buf bytes.Buffer
	buf.WriteString("errors: ")
	if len(t.errors) > 0 {
		buf.WriteString("\n")
		for i, err := range t.errors {
			buf.WriteString(fmt.Sprintf("\t%d. %s", i, err.Error()))
			buf.WriteString("\n")
		}
	} else {
		buf.WriteString("none")
	}
	return buf.String()
}
