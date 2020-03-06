package lexer

import (
	"github.com/sonirico/datetoken.go/token"
	"strings"
)

type Lexer struct {
	payload        string
	limit          int64
	currentChar    byte
	currentPointer int64
	nextPointer    int64
}

func NewLexer(payload string) *Lexer {
	cleanPayload := strings.TrimSpace(payload)
	lexer := &Lexer{
		payload:        cleanPayload,
		limit:          int64(len(cleanPayload)),
		currentPointer: 0,
		nextPointer:    0,
	}
	lexer.readChar()
	return lexer
}

func (l *Lexer) readChar() {
	if l.nextPointer < l.limit {
		l.currentChar = l.payload[l.nextPointer]
	} else {
		l.currentChar = 0
	}
	l.currentPointer = l.nextPointer
	l.nextPointer += 1
}

func (l *Lexer) peekChar() byte {
	if l.nextPointer < l.limit {
		return l.payload[l.nextPointer]
	}
	return 0
}

func (l *Lexer) readWord() string {
	pos := l.currentPointer
	for isLetter(l.currentChar) {
		l.readChar()
	}
	return l.payload[pos:l.currentPointer]
}

func (l *Lexer) readNumber() string {
	pos := l.currentPointer
	for isDigit(l.currentChar) {
		l.readChar()
	}
	return l.payload[pos:l.currentPointer]
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	switch l.currentChar {
	case 'n':
		word := l.readWord()
		if "now" == word {
			tok = newToken(token.Start, "now")
		} else {
			tok = newToken(token.Illegal, word)
		}
		return tok
	case '+':
		tok = newToken(token.Plus, "+")
	case '-':
		tok = newToken(token.Minus, "-")
	case '/':
		tok = newToken(token.SnapStart, "/")
	case '@':
		tok = newToken(token.SnapEnd, "@")
	case 0:
		tok = newToken(token.End, "")
	default:
		if isDigit(l.currentChar) {
			number := l.readNumber()
			tok = newToken(token.Number, number)
			return tok
		} else {
			schar := string(l.currentChar)
			tok = newToken(token.LookupKeyword(schar), schar)
		}
	}
	l.readChar()
	return tok
}

func isLetter(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || char == '_'
}

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func newToken(tt token.TokenType, literal string) token.Token {
	return token.NewToken(tt, literal)
}
