package models

import "time"

type Token struct {
	Literal string
	Date    time.Time
	Errors  []error
}
