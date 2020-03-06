package token

type TokenType string

const (
	Start     TokenType = "start"
	End                 = "end"
	SnapStart           = "snapstart/"
	SnapEnd             = "snapend@"
	Plus                = "plus+"
	Minus               = "minus-"
	Number              = "num"
	Unit                = "unit"
	Illegal             = "ill"
)

var keywords = map[string]TokenType{
	"+": Plus,
	"-": Minus,
	"/": SnapStart,
	"@": SnapEnd,
	"s": Unit,
	"m": Unit,
	"h": Unit,
	"d": Unit,
	"w": Unit,
	"M": Unit,
	"Y": Unit,
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
