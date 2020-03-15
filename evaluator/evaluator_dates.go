package evaluator

func (e *Evaluator) snapStartOfSecond() {
	e.current = BeginningOfSecond(e.current)
}

// BeginningOfMinute beginning of minute
func (e *Evaluator) snapStartOfMinute() {
	e.current = BeginningOfMinute(e.current)
}

// BeginningOfHour beginning of hour
func (e *Evaluator) snapStartOfHour() {
	e.current = BeginningOfHour(e.current)
}

// BeginningOfDay beginning of day
func (e *Evaluator) snapStartOfDay() {
	e.current = BeginningOfDay(e.current)
}

// BeginningOfWeek beginning of week
func (e *Evaluator) snapStartOfWeek() {
	e.current = BeginningOfWeek(e.current, e.weekStartDay)
}

// BeginningOfMonth beginning of month
func (e *Evaluator) snapStartOfMonth() {
	e.current = BeginningOfMonth(e.current)
}

// BeginningOfYear BeginningOfYear beginning of year
func (e *Evaluator) snapStartOfYear() {
	e.current = BeginningOfYear(e.current)
}

// EndOfMinute end of minute
func (e *Evaluator) snapEndOfMinute() {
	e.current = EndOfMinute(e.current)
}

// EndOfHour end of hour
func (e *Evaluator) snapEndOfHour() {
	e.current = EndOfHour(e.current)
}

// EndOfDay end of day
func (e *Evaluator) snapEndOfDay() {
	e.current = EndOfDay(e.current)
}

// EndOfWeek end of week
func (e *Evaluator) snapEndOfWeek() {
	e.current = EndOfWeek(e.current, e.weekStartDay)
}

// EndOfMonth end of month
func (e *Evaluator) snapEndOfMonth() {
	e.current = EndOfMonth(e.current)
}

// EndOfYear end of year
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
