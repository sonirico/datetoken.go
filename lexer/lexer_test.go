package lexer

import (
	"fmt"
	"testing"

	"github.com/sonirico/datetoken.go/token"
)

type expectedResult struct {
	Type    token.TokenType
	Literal string
}

func testToken(t *testing.T, payload string, expected []expectedResult) {
	t.Helper()

	lexer := New(payload)

	for _, expectedToken := range expected {
		actualToken := lexer.NextToken()
		if expectedToken.Type != actualToken.Type {
			t.Fatalf("unexpected token type. want '%v', have '%v'", expectedToken, actualToken)
		}
		if expectedToken.Literal != actualToken.Literal {
			t.Fatalf("unexpected token literal. want '%v', have '%v'", expectedToken, actualToken)
		}
		fmt.Printf(fmt.Sprintf("token ok: %v\n", actualToken))
	}

}
func TestLexerWithNow(t *testing.T) {
	input := `
		now-s+2m_-3h+234d-w+M-Y/M@w
	`
	expected := []expectedResult{
		{token.Start, "now"},
		{token.Minus, "-"},
		{token.Unit, "s"},
		{token.Plus, "+"},
		{token.Number, "2"},
		{token.Unit, "m"},
		{token.Illegal, "_"},
		{token.Minus, "-"},
		{token.Number, "3"},
		{token.Unit, "h"},
		{token.Plus, "+"},
		{token.Number, "234"},
		{token.Unit, "d"},
		{token.Minus, "-"},
		{token.Unit, "w"},
		{token.Plus, "+"},
		{token.Unit, "M"},
		{token.Minus, "-"},
		{token.Unit, "Y"},
		{token.SnapStart, "/"},
		{token.Unit, "M"},
		{token.SnapEnd, "@"},
		{token.Unit, "w"},
		{token.End, ""},
	}

	testToken(t, input, expected)
}

func TestLexerWithoutNow(t *testing.T) {
	input := `
		-s+2m_-3h+234d-w+M-Y/M@w
	`
	expected := []expectedResult{
		{token.Minus, "-"},
		{token.Unit, "s"},
		{token.Plus, "+"},
		{token.Number, "2"},
		{token.Unit, "m"},
		{token.Illegal, "_"},
		{token.Minus, "-"},
		{token.Number, "3"},
		{token.Unit, "h"},
		{token.Plus, "+"},
		{token.Number, "234"},
		{token.Unit, "d"},
		{token.Minus, "-"},
		{token.Unit, "w"},
		{token.Plus, "+"},
		{token.Unit, "M"},
		{token.Minus, "-"},
		{token.Unit, "Y"},
		{token.SnapStart, "/"},
		{token.Unit, "M"},
		{token.SnapEnd, "@"},
		{token.Unit, "w"},
		{token.End, ""},
	}

	testToken(t, input, expected)
}
