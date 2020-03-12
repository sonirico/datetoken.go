package evaluator

func (e *evaluator) snapStartOfSecond() {
	e.current = BeginningOfSecond(e.current)
}

// BeginningOfMinute beginning of minute
func (e *evaluator) snapStartOfMinute() {
	e.current = BeginningOfMinute(e.current)
}

// BeginningOfHour beginning of hour
func (e *evaluator) snapStartOfHour() {
	e.current = BeginningOfHour(e.current)
}

// BeginningOfDay beginning of day
func (e *evaluator) snapStartOfDay() {
	e.current = BeginningOfDay(e.current)
}

// BeginningOfWeek beginning of week
func (e *evaluator) snapStartOfWeek() {
	e.current = BeginningOfWeek(e.current, 0)
}

// BeginningOfMonth beginning of month
func (e *evaluator) snapStartOfMonth() {
	e.current = BeginningOfMonth(e.current)
}

// BeginningOfYear BeginningOfYear beginning of year
func (e *evaluator) snapStartOfYear() {
	e.current = BeginningOfYear(e.current)
}

// EndOfMinute end of minute
func (e *evaluator) snapEndOfMinute() {
	e.current = EndOfMinute(e.current)
}

// EndOfHour end of hour
func (e *evaluator) snapEndOfHour() {
	e.current = EndOfHour(e.current)
}

// EndOfDay end of day
func (e *evaluator) snapEndOfDay() {
	e.current = EndOfDay(e.current)
}

// EndOfWeek end of week
func (e *evaluator) snapEndOfWeek() {
	e.current = EndOfWeek(e.current, 0)
}

// EndOfMonth end of month
func (e *evaluator) snapEndOfMonth() {
	e.current = EndOfMonth(e.current)
}

// EndOfYear end of year
func (e *evaluator) snapEndOfYear() {
	e.current = EndOfYear(e.current)
}
