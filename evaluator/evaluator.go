package evaluator

import (
	"errors"
	"time"

	"github.com/sonirico/datetoken.go/ast"
	"github.com/sonirico/datetoken.go/lexer"
	"github.com/sonirico/datetoken.go/parser"
	"github.com/sonirico/datetoken.go/token"
)

type evalArithmeticFunc func(Unit) time.Time
type evalSnapFunc func(Snap) time.Time

type evalConfig struct {
	WeekStartDay time.Weekday
	TimeLocation *time.Location
}

type evaluator struct {
	time.Time

	config *evalConfig

	initial *time.Time
	tmp     *time.Time

	arithmeticRegistry map[Unit]evalArithmeticFunc
	snapRegistry       map[Snap]evalSnapFunc
}

func newEvaluator() *evaluator {
	evaluator := &evaluator{
		initial: nil,
		tmp:     nil,
		config: &evalConfig{
			WeekStartDay: time.Sunday,
		},
		arithmeticRegistry: make(map[Unit]evalArithmeticFunc),
		snapRegistry:       make(map[Snap]evalSnapFunc),
	}
	return evaluator
}

func (e *evaluator) registerSnapFunc(snap Snap, fn evalSnapFunc) {
	e.snapRegistry[snap] = fn
}

func (e *evaluator) registerArithmeticFunc(unit Unit, fn evalArithmeticFunc) {
	e.arithmeticRegistry[unit] = fn
}

// Override initial node value
func (e *evaluator) setInitial(date time.Time) {
	if e.initial != nil {
		return
	}
	e.initial = &date
	if e.tmp == nil {
		e.tmp = &date
	}
}

func (e *evaluator) evalValueNode(node *ast.ValueNode) {
	switch node.Literal() {
	case Now:
	default:
		e.setInitial(time.Now())
	}
}

func (e *evaluator) evalArithmeticNode(node *ast.ArithmeticNode) {

}

func (e *evaluator) evalStartSnap(node *ast.SnapNode) {
	var newTmp time.Time
	switch node.Unit {
	case "s":
		newTmp = e.beginningOfSecond()
	case "m":
		newTmp = e.beginningOfMinute()
	case "h":
		newTmp = e.beginningOfHour()
	case "d":
		newTmp = e.beginningOfDay()
	case "w":
		newTmp = e.beginningOfWeek()
	case "M":
		newTmp = e.beginningOfMonth()
	case "Y":
		newTmp = e.beginningOfYear()
	}
	e.tmp = &newTmp
}

func (e *evaluator) evalEndSnap(node *ast.SnapNode) {
	var newTmp time.Time
	switch node.Unit {
	case "s":
		newTmp = *e.initial
	case "m":
		newTmp = e.endOfMinute()
	case "h":
		newTmp = e.endOfHour()
	case "d":
		newTmp = e.endOfDay()
	case "w":
		newTmp = e.endOfWeek()
	case "M":
		newTmp = e.endOfMonth()
	case "Y":
		newTmp = e.endOfYear()
	}
	e.tmp = &newTmp
}

func (e *evaluator) evalSnapNode(node *ast.SnapNode) {
	switch node.Token.Type {
	case token.SnapStart:
		e.evalStartSnap(node)
	case token.SnapEnd:
		e.evalEndSnap(node)
	}
}

func (e *evaluator) evalNode(node ast.Node) {
	switch nod := node.(type) {
	case *ast.ValueNode:
		e.evalValueNode(nod)
	case *ast.ArithmeticNode:
		e.evalArithmeticNode(nod)
	case *ast.SnapNode:
		e.evalSnapNode(nod)
	}
}

func (e *evaluator) Eval(datetoken string) (time.Time, error) {
	lexer := lexer.New(datetoken)
	parser := parser.New(lexer)
	astRoot := parser.Parse()
	if len(parser.Errors()) > 1 {
		return time.Now(), errors.New("parser errors")
	}
	for _, node := range astRoot.Nodes {
		e.evalNode(node)
	}
	return *e.tmp, nil
}

func Eval(payload string) (time.Time, error) {
	evaluatorImpl := &evaluator{}
	return evaluatorImpl.Eval(payload)
}
