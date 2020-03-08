package evaluator

type Value string

type Unit string

type Snap string

const (
	Second Unit = "s"
	Minute Unit = "m"
	Hour   Unit = "h"
	Day    Unit = "d"
	Week   Unit = "w"
	Month  Unit = "M"
	Year   Unit = "Y"
)

const (
	Start Snap = "/"
	End   Snap = "@"
)

const (
	Now = "now"
)
