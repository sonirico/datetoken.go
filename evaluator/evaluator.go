package evaluator

import (
	"errors"
	"time"

	"github.com/sonirico/datetoken.go/ast"
	"github.com/sonirico/datetoken.go/lexer"
	"github.com/sonirico/datetoken.go/parser"
	"github.com/sonirico/datetoken.go/token"
)

type evalConfig struct {
	WeekStartDay time.Weekday
	TimeLocation *time.Location
}

type evaluator struct {
	time.Time

	config *evalConfig

	current time.Time
	timeSet bool
}

func newEvaluator() *evaluator {
	evaluator := &evaluator{
		config: &evalConfig{
			WeekStartDay: time.Sunday,
		},
		timeSet: false,
	}
	return evaluator
}

// Override initial node value
func (e *evaluator) setInitial(date time.Time) {
	if e.timeSet {
		return
	}
	e.current = date
	e.timeSet = true
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
	case Second:
		e.addSeconds(amount)
	case Minute:
		e.addMinutes(amount)
	case Hour:
		e.addHours(amount)
	case Day:
		e.addDays(amount)
	case Week:
		e.addWeeks(amount)
	case Month:
		e.addMonths(amount)
	case Year:
		e.addYears(amount)
	}
}

func (e *evaluator) evalStartSnap(node *ast.SnapNode) {
	switch node.Unit {
	// time units
	case Minute:
		e.snapStartOfMinute()
	case Hour:
		e.snapStartOfHour()
	case Day:
		e.snapStartOfDay()
	case Week:
		e.snapStartOfWeek()
	case Month:
		e.snapStartOfMonth()
	case Year:
		e.snapStartOfYear()
	// weekdays
	case Monday:
		e.previousMonday()
	case Tuesday:
		e.previousTuesday()
	case Wednesday:
		e.previousWednesday()
	case Thursday:
		e.previousThursday()
	case Friday:
		e.previousFriday()
	case Saturday:
		e.previousSaturday()
	case Sunday:
		e.previousSunday()
	}
}

func (e *evaluator) evalEndSnap(node *ast.SnapNode) {
	switch node.Unit {
	// time unit
	case Minute:
		e.snapEndOfMinute()
	case Hour:
		e.snapEndOfHour()
	case Day:
		e.snapEndOfDay()
	case Week:
		e.snapEndOfWeek()
	case Month:
		e.snapEndOfMonth()
	case Year:
		e.snapEndOfYear()
	// weekdays
	case Monday:
		e.nextMonday()
	case Tuesday:
		e.nextTuesday()
	case Wednesday:
		e.nextWednesday()
	case Thursday:
		e.nextThursday()
	case Friday:
		e.nextFriday()
	case Saturday:
		e.nextSaturday()
	case Sunday:
		e.nextSunday()
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
