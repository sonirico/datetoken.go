package evaluator

import (
	"github.com/sonirico/datetoken.go/models"
	"time"

	"github.com/sonirico/datetoken.go/ast"
	"github.com/sonirico/datetoken.go/lexer"
	"github.com/sonirico/datetoken.go/parser"
	"github.com/sonirico/datetoken.go/token"
)

// Evaluator takes a token payload to eval. Handles lexing and parsing too.
type Evaluator struct {
	time.Time

	weekStartDay time.Weekday
	tz           *time.Location

	current time.Time
	timeSet bool
}

// New returns a new instance of Evaluator
func New() *Evaluator {
	return &Evaluator{
		weekStartDay: time.Sunday,
		tz:           time.UTC,
		timeSet:      false,
	}
}

// SetTZ allows to configure a different time.Location rather than UTC
func (e *Evaluator) SetTZ(tz *time.Location) {
	e.tz = tz
}

// SetWeeksStartDay allows to configure a different time.WeekDay rather than Sunday
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
	case now:
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
	case second:
		e.addSeconds(amount)
	case minute:
		e.addMinutes(amount)
	case hour:
		e.addHours(amount)
	case day:
		e.addDays(amount)
	case week:
		e.addWeeks(amount)
	case month:
		e.addMonths(amount)
	case year:
		e.addYears(amount)
	}
}

func (e *Evaluator) evalStartSnap(node *ast.SnapNode) {
	switch node.Unit {
	// time units
	case minute:
		e.snapStartOfMinute()
	case hour:
		e.snapStartOfHour()
	case day:
		e.snapStartOfDay()
	case week:
		e.snapStartOfWeek()
	case businessWeek:
		e.snapStartOfBusinessWeek()
	case month:
		e.snapStartOfMonth()
	case year:
		e.snapStartOfYear()
	// weekdays
	case monday:
		e.previousMonday()
	case tuesday:
		e.previousTuesday()
	case wednesday:
		e.previousWednesday()
	case thursday:
		e.previousThursday()
	case friday:
		e.previousFriday()
	case saturday:
		e.previousSaturday()
	case sunday:
		e.previousSunday()
	}
}

func (e *Evaluator) evalEndSnap(node *ast.SnapNode) {
	switch node.Unit {
	// time unit
	case minute:
		e.snapEndOfMinute()
	case hour:
		e.snapEndOfHour()
	case day:
		e.snapEndOfDay()
	case week:
		e.snapEndOfWeek()
	case businessWeek:
		e.snapEndOfBusinessWeek()
	case month:
		e.snapEndOfMonth()
	case year:
		e.snapEndOfYear()
	// weekdays
	case monday:
		e.nextMonday()
	case tuesday:
		e.nextTuesday()
	case wednesday:
		e.nextWednesday()
	case thursday:
		e.nextThursday()
	case friday:
		e.nextFriday()
	case saturday:
		e.nextSaturday()
	case sunday:
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

// Eval takes a token payload and evaluates it. If correct, it returns the time.Time. Otherwise, the current
// timestamp is returned along with errors
func (e *Evaluator) Eval(payload string) (time.Time, error) {
	lex := lexer.New(payload)
	par := parser.New(lex)
	astRoot := par.Parse()
	if len(astRoot.Nodes) < 1 {
		return time.Now(), models.ErrEmptyToken
	}
	if len(par.Errors()) > 0 {
		return time.Now(), models.NewInvalidTokenError(payload, par.Errors())
	}
	for _, node := range astRoot.Nodes {
		e.evalNode(node)
	}
	return e.current, nil
}
