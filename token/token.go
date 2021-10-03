package token

// Type represent the kind of possible token categories
type Type string

const (
	// Start is the start mark of the tokenization
	Start Type = "start"
	// End is the end mark of the tokenization
	End = "end"
	// SnapStart is the token mapped to snap dates to start of an unit
	SnapStart = "/"
	// SnapEnd is the token mapped to snap dates to end of an unit
	SnapEnd = "@"
	// Plus is the token mapped to add time units to a date
	Plus = "+"
	// Minus is the token mapped to subtract time units to a date
	Minus = "-"
	// Number is the token for any number in a token
	Number = "num"
	// Unit is the token for any time unit in a token
	Unit = "unit"
	// Wd is the token for any weekday
	Wd = "wd"
	// Illegal represents any not supported nor understood value
	Illegal = "ill"
)

var keywords = map[string]Type{
	"+":   Plus,
	"-":   Minus,
	"/":   SnapStart,
	"@":   SnapEnd,
	"s":   Unit,
	"m":   Unit,
	"h":   Unit,
	"d":   Unit,
	"w":   Unit,
	"bw":  Unit,
	"M":   Unit,
	"Y":   Unit,
	"mon": Wd,
	"tue": Wd,
	"wed": Wd,
	"thu": Wd,
	"fri": Wd,
	"sat": Wd,
	"sun": Wd,
	"Q":   Unit,
	"Q1":  Unit,
	"Q2":  Unit,
	"Q3":  Unit,
	"Q4":  Unit,
}

// LookupKeyword will return the associated token type for a given token literal. If no one is found, the literal
// is considered to be illegal
func LookupKeyword(key string) Type {
	if tt, ok := keywords[key]; ok {
		return tt
	}
	return Illegal
}

// Token represents the minimal piece of information relevant to datetoken
type Token struct {
	// Type indicates the token type
	Type Type
	// Literal stores the literal string from the input source
	Literal string
}

// NewToken returns a new instance of Token
func NewToken(tt Type, literal string) Token {
	return Token{Type: tt, Literal: literal}
}
