package evaluator

func (e *Evaluator) snapStartOfSecond() {
	e.current = BeginningOfSecond(e.current)
}

func (e *Evaluator) snapStartOfMinute() {
	e.current = BeginningOfMinute(e.current)
}

func (e *Evaluator) snapStartOfHour() {
	e.current = BeginningOfHour(e.current)
}

func (e *Evaluator) snapStartOfDay() {
	e.current = BeginningOfDay(e.current)
}

func (e *Evaluator) snapStartOfWeek() {
	e.current = BeginningOfWeek(e.current, e.weekStartDay)
}

func (e *Evaluator) snapStartOfBusinessWeek() {
	e.snapStartOfWeek()
}

func (e *Evaluator) snapStartOfMonth() {
	e.current = BeginningOfMonth(e.current)
}

func (e *Evaluator) snapStartOfYear() {
	e.current = BeginningOfYear(e.current)
}

func (e *Evaluator) snapEndOfMinute() {
	e.current = EndOfMinute(e.current)
}

func (e *Evaluator) snapEndOfHour() {
	e.current = EndOfHour(e.current)
}

func (e *Evaluator) snapEndOfDay() {
	e.current = EndOfDay(e.current)
}

func (e *Evaluator) snapEndOfWeek() {
	e.current = EndOfWeek(e.current, e.weekStartDay)
}

func (e *Evaluator) snapEndOfBusinessWeek() {
	e.current = EndOfBusinessWeek(e.current, e.weekStartDay)
}

func (e *Evaluator) snapEndOfMonth() {
	e.current = EndOfMonth(e.current)
}

func (e *Evaluator) snapEndOfYear() {
	e.current = EndOfYear(e.current)
}

func (e *Evaluator) addSeconds(amount int) {
	e.current = AddSeconds(e.current, amount)
}

func (e *Evaluator) addMinutes(amount int) {
	e.current = AddMinutes(e.current, amount)
}

func (e *Evaluator) addHours(amount int) {
	e.current = AddHours(e.current, amount)
}

func (e *Evaluator) addDays(amount int) {
	e.current = AddDays(e.current, amount)
}

func (e *Evaluator) addWeeks(amount int) {
	e.current = AddWeeks(e.current, amount)
}

func (e *Evaluator) addMonths(amount int) {
	e.current = AddMonths(e.current, amount)
}

func (e *Evaluator) addYears(amount int) {
	e.current = AddYears(e.current, amount)
}

func (e *Evaluator) previousMonday() {
	e.current = PreviousMonday(e.current)
}

func (e *Evaluator) previousTuesday() {
	e.current = PreviousTuesday(e.current)
}

func (e *Evaluator) previousWednesday() {
	e.current = PreviousWednesday(e.current)
}

func (e *Evaluator) previousThursday() {
	e.current = PreviousThursday(e.current)
}

func (e *Evaluator) previousFriday() {
	e.current = PreviousFriday(e.current)
}

func (e *Evaluator) previousSaturday() {
	e.current = PreviousSaturday(e.current)
}

func (e *Evaluator) previousSunday() {
	e.current = PreviousSunday(e.current)
}

func (e *Evaluator) nextMonday() {
	e.current = NextMonday(e.current)
}

func (e *Evaluator) nextTuesday() {
	e.current = NextTuesday(e.current)
}

func (e *Evaluator) nextWednesday() {
	e.current = NextWednesday(e.current)
}

func (e *Evaluator) nextThursday() {
	e.current = NextThursday(e.current)
}

func (e *Evaluator) nextFriday() {
	e.current = NextFriday(e.current)
}

func (e *Evaluator) nextSaturday() {
	e.current = NextSaturday(e.current)
}

func (e *Evaluator) nextSunday() {
	e.current = NextSunday(e.current)
}
