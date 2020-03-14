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

	current time.Time
	timeSet bool

	arithmeticRegistry map[Unit]evalArithmeticFunc
	snapRegistry       map[Snap]evalSnapFunc
}

func newEvaluator() *evaluator {
	evaluator := &evaluator{
		config: &evalConfig{
			WeekStartDay: time.Sunday,
		},
		timeSet:            false,
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
	if e.timeSet {
		return
	}
	e.current = date
}

func (e *evaluator) evalValueNode(node *ast.ValueNode) {
	switch node.Literal() {
	case Now:
	default:
		e.setInitial(time.Now())
	}
}

func (e *evaluator) evalArithmeticNode(node *ast.ArithmeticNode) {
	amount := int(node.Amount)
	if token.Minus == node.Sign {
		amount = -amount
	}

	switch node.Unit {
	case "s":
		e.addSeconds(amount)
	case "m":
		e.addMinutes(amount)
	case "h":
		e.addHours(amount)
	case "d":
		e.addDays(amount)
	case "w":
		e.addWeeks(amount)
	case "M":
		e.addMonths(amount)
	case "Y":
		e.addYears(amount)
	}
}

func (e *evaluator) evalStartSnap(node *ast.SnapNode) {
	switch node.Unit {
	case "m":
		e.snapStartOfMinute()
	case "h":
		e.snapStartOfHour()
	case "d":
		e.snapStartOfDay()
	case "w":
		e.snapStartOfWeek()
	case "M":
		e.snapStartOfMonth()
	case "Y":
		e.snapStartOfYear()
	}
}

func (e *evaluator) evalEndSnap(node *ast.SnapNode) {
	switch node.Unit {
	case "m":
		e.snapEndOfMinute()
	case "h":
		e.snapEndOfHour()
	case "d":
		e.snapEndOfDay()
	case "w":
		e.snapEndOfWeek()
	case "M":
		e.snapEndOfMonth()
	case "Y":
		e.snapEndOfYear()
	}
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

func (e *evaluator) Eval(payload string) (time.Time, error) {
	lexer := lexer.New(payload)
	parser := parser.New(lexer)
	astRoot := parser.Parse()
	if len(parser.Errors()) > 1 {
		return time.Now(), errors.New("parser errors")
	}
	for _, node := range astRoot.Nodes {
		e.evalNode(node)
	}
	return e.current, nil
}

func Eval(payload string) (time.Time, error) {
	evaluatorImpl := &evaluator{}
	return evaluatorImpl.Eval(payload)
}
