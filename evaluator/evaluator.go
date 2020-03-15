package evaluator

import (
	"errors"
	"time"

	"github.com/sonirico/datetoken.go/ast"
	"github.com/sonirico/datetoken.go/lexer"
	"github.com/sonirico/datetoken.go/parser"
	"github.com/sonirico/datetoken.go/token"
)

type Evaluator struct {
	time.Time

	weekStartDay time.Weekday
	tz           *time.Location

	current time.Time
	timeSet bool
}

func New() *Evaluator {
	return &Evaluator{
		weekStartDay: time.Sunday,
		tz:           time.UTC,
		timeSet:      false,
	}
}

func (e *Evaluator) SetTZ(tz *time.Location) {
	e.tz = tz
}

func (e *Evaluator) SetWeeksStartDay(wd time.Weekday) {
	e.weekStartDay = wd
}

// Override initial node value
func (e *Evaluator) setInitial(date time.Time) {
	if e.timeSet {
		return
	}
	e.current = date
	if e.tz != nil {
		e.current = e.current.In(e.tz)
	}
	e.timeSet = true
}

func (e *Evaluator) evalValueNode(node *ast.ValueNode) {
	switch node.Literal() {
	case Now:
		fallthrough
	default:
		e.setInitial(time.Now())
	}
}

func (e *Evaluator) evalArithmeticNode(node *ast.ArithmeticNode) {
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

func (e *Evaluator) evalStartSnap(node *ast.SnapNode) {
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

func (e *Evaluator) evalEndSnap(node *ast.SnapNode) {
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

func (e *Evaluator) evalSnapNode(node *ast.SnapNode) {
	switch node.Token.Type {
	case token.SnapStart:
		e.evalStartSnap(node)
	case token.SnapEnd:
		e.evalEndSnap(node)
	}
}

func (e *Evaluator) evalNode(node ast.Node) {
	switch nod := node.(type) {
	case *ast.ValueNode:
		e.evalValueNode(nod)
	case *ast.ArithmeticNode:
		e.evalArithmeticNode(nod)
	case *ast.SnapNode:
		e.evalSnapNode(nod)
	}
}

func (e *Evaluator) Eval(payload string) (time.Time, error) {
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
