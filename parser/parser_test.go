package parser

import (
	"testing"

	"github.com/sonirico/datetoken.go/ast"
	"github.com/sonirico/datetoken.go/lexer"
)

func newParser(t *testing.T, payload string) *ast.RootNode {
	lexer := lexer.New(payload)
	parser := New(lexer)
	root := parser.Parse()
	if len(parser.Errors()) > 0 {
		t.Errorf("parser had %d errors", len(parser.Errors()))
		for _, msg := range parser.Errors() {
			t.Errorf("parser error: %s", msg)
		}
		t.FailNow()
	}
	return root
}

func testValNode(t *testing.T, node ast.Node, expectedLiteral string) bool {
	valnode, ok := node.(*ast.ValueNode)
	if !ok {
		t.Fatalf("unexpected node type. want ValueNode, have %v(%T)",
			valnode, valnode)
		return false
	}
	if valnode.Literal() != "now" {
		t.Fatalf("valuenode. unexpected literal. want %s, have %s",
			"now", valnode.Literal())
		return false
	}
	return true
}

func testArithmeticNode(t *testing.T, node ast.Node, amount int64, unit, sign string) bool {
	arnode, ok := node.(*ast.ArithmeticNode)
	if !ok {
		t.Fatalf("unexpected node type. want ArithmeticNode, have %v(%T)",
			arnode, arnode)
		return false
	}
	if arnode.Amount != amount {
		t.Fatalf("arnode. unexpected amount. want %d, have %d",
			amount, arnode.Amount)
		return false
	}
	if arnode.Unit != unit {
		t.Fatalf("arnode. unexpected unit. want %s, have %s",
			unit, arnode.Unit)
		return false
	}
	if arnode.Sign != sign {
		t.Fatalf("arnode. unexpected sign symbol. want %s, have %s",
			sign, arnode.Sign)
	}
	return true
}

func testSnapNode(t *testing.T, node ast.Node, op, unit string) bool {
	snapnode, ok := node.(*ast.SnapNode)
	if !ok {
		t.Fatalf("unexpected node type. want SnapNode, have %v(%T)",
			snapnode, snapnode)
		return false
	}
	if snapnode.Operator != op {
		t.Fatalf("snapnode. unexpected operator. want %s, have %s",
			op, snapnode.Operator)
		return false
	}
	if snapnode.Unit != unit {
		t.Fatalf("snapnode. unexpected unit. want %s, have %s",
			unit, snapnode.Unit)
		return false
	}
	return true
}

func TestParser_SnapStart(t *testing.T) {
	payload := "now/d"
	root := newParser(t, payload)
	if len(root.Nodes) < 2 {
		t.Fatalf("empty node set. expected some")
	}
	testValNode(t, root.Nodes[0], "now")
	testSnapNode(t, root.Nodes[1], "/", "d")
}

func TestParser_SnapEnd(t *testing.T) {
	payload := "now@d"
	root := newParser(t, payload)
	if len(root.Nodes) < 2 {
		t.Fatalf("empty node set. expected some")
	}
	testValNode(t, root.Nodes[0], "now")
	testSnapNode(t, root.Nodes[1], "@", "d")
}

func TestParser_SnapEnd_Weekday(t *testing.T) {
	payload := "now@thu"
	root := newParser(t, payload)
	if len(root.Nodes) < 2 {
		t.Fatalf("empty node set. expected some")
	}
	testValNode(t, root.Nodes[0], "now")
	testSnapNode(t, root.Nodes[1], "@", "thu")
}

func TestParser_ArithmeticPlus(t *testing.T) {
	payload := "now+01d"
	root := newParser(t, payload)
	if len(root.Nodes) < 1 {
		t.Fatalf("empty node set. expected some")
	}
	testValNode(t, root.Nodes[0], "now")
	testArithmeticNode(t, root.Nodes[1], 1, "d", "+")
}

func TestParser_ArithmeticPlusImplicitAmount(t *testing.T) {
	payload := "now+d"
	root := newParser(t, payload)
	if len(root.Nodes) < 1 {
		t.Fatalf("empty node set. expected some")
	}
	testValNode(t, root.Nodes[0], "now")
	testArithmeticNode(t, root.Nodes[1], 1, "d", "+")
}
func TestParser_ArithmeticMinus(t *testing.T) {
	payload := "now-23d"
	root := newParser(t, payload)
	if len(root.Nodes) < 1 {
		t.Fatalf("empty node set. expected some")
	}
	testValNode(t, root.Nodes[0], "now")
	testArithmeticNode(t, root.Nodes[1], 23, "d", "-")
}

func TestParser_ArithmeticMinusImplicitAmount(t *testing.T) {
	payload := "now-d"
	root := newParser(t, payload)
	if len(root.Nodes) < 1 {
		t.Fatalf("empty node set. expected some")
	}
	testValNode(t, root.Nodes[0], "now")
	testArithmeticNode(t, root.Nodes[1], 1, "d", "-")
}
