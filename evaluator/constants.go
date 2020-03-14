package evaluator

import "time"

const (
	Second string = "s"
	Minute        = "m"
	Hour          = "h"
	Day           = "d"
	Week          = "w"
	Month         = "M"
	Year          = "Y"

	Monday    = "mon"
	Tuesday   = "tue"
	Wednesday = "wed"
	Thursday  = "thu"
	Friday    = "fri"
	Saturday  = "sat"
	Sunday    = "sun"
)

const (
	Start string = "/"
	End   string = "@"
)

const (
	Now = "now"
)

type valueResolver map[string]func(string) time.Time

func (vr valueResolver) register(token string, fn func(string) time.Time) {
	vr[token] = fn
}
