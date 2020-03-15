package token

type TokenType string

const (
	Start     TokenType = "start"
	End                 = "end"
	SnapStart           = "/"
	SnapEnd             = "@"
	Plus                = "+"
	Minus               = "-"
	Number              = "num"
	Unit                = "unit"
	Illegal             = "ill"
	Wd                  = "wd"
)

var keywords = map[string]TokenType{
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
}

func LookupKeyword(key string) TokenType {
	if tt, ok := keywords[key]; ok {
		return tt
	}
	return Illegal
}

type Token struct {
	Type    TokenType
	Literal string
}

func NewToken(tt TokenType, literal string) Token {
	return Token{Type: tt, Literal: literal}
}
