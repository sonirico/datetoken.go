package models

import (
	"bytes"
	"fmt"
	"time"
)

type Token struct {
	literal string    `json:"literal"`
	date    time.Time `json:"datetime"`
	errors  []error   `json:"errors"`
}

func NewToken(payload string, date time.Time) Token {
	return Token{
		literal: payload,
		date:    date,
		errors:  make([]error, 0),
	}
}

func (t Token) AddError(err error) {
	t.errors = append(t.errors, err)
}

func (t Token) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("payload: '%s'", t.literal))
	buf.WriteString("\n")
	buf.WriteString(fmt.Sprintf("date: '%s'", t.date.String()))
	buf.WriteString("\n")
	buf.WriteString(t.ErrorString())
	return buf.String()
}

func (t Token) Errors() []error {
	return t.errors
}

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
