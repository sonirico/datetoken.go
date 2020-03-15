package lexer

import (
	"fmt"
	"testing"

	"github.com/sonirico/datetoken.go/token"
)

type expectedResult struct {
	Type    token.Type
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
func TestLexer_WithNow(t *testing.T) {
	input := `
		now-s+2m_-3h+234d-w+M-Y/M@w@bw
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
		{token.SnapEnd, "@"},
		{token.Unit, "bw"},
		{token.End, ""},
	}

	testToken(t, input, expected)
}

func TestLexer_WithoutNow(t *testing.T) {
	input := `
		-s+2m_-3h+234d-w+M-Y/M@w/bw
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
		{token.SnapStart, "/"},
		{token.Unit, "bw"},
		{token.End, ""},
	}

	testToken(t, input, expected)
}

func TestLexer_PreviousWeekdays(t *testing.T) {
	input := "now/mon/tue/wed/thu/fri/sat/sun"
	expected := []expectedResult{
		{token.Start, "now"},
		{token.SnapStart, "/"},
		{token.Wd, "mon"},
		{token.SnapStart, "/"},
		{token.Wd, "tue"},
		{token.SnapStart, "/"},
		{token.Wd, "wed"},
		{token.SnapStart, "/"},
		{token.Wd, "thu"},
		{token.SnapStart, "/"},
		{token.Wd, "fri"},
		{token.SnapStart, "/"},
		{token.Wd, "sat"},
		{token.SnapStart, "/"},
		{token.Wd, "sun"},
		{token.End, ""},
	}

	testToken(t, input, expected)
}
