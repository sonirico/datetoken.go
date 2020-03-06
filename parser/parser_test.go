package parser

import (
	"testing"

	"github.com/sonirico/datetoken.go/ast"
	"github.com/sonirico/datetoken.go/lexer"
)

func newParser(payload string) *ast.RootNode {
	lexer := lexer.NewLexer(payload)
	parser := NewParser(lexer)
	root := parser.Parse()
	return root
}

func TestParser_SnapStart(t *testing.T) {
	payload := "now/d"
	root := newParser(payload)
	expected := []struct {
		Literal string
		Type    ast.NodeType
	}{}

	if len(root.Nodes) != len(expected) {
		t.Errorf("unexpected number of ast nodes. want %d, have %d",
			len(expected), len(root.Nodes))
	}

	for index, node := range root.Nodes {
		expectedNode := expected[index]
		if node.Type() != expectedNode.Type {
			t.Errorf("unexpected node type. want %s, have %s",
				expectedNode.Type, node.Type())
		}

		if node.Inspect() != expectedNode.Literal {
			t.Errorf("unexpected node literal. want %s, have %s",
				expectedNode.Type, node.Type())
		}
	}
}
