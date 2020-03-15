package datetoken

import "time"

// Config provides an interface to configure how tokens are going to be evaluated
type Config interface {
	// Tz should provide the time zone location of the client
	Tz() *time.Location
	// WeeksStartAt should yield, depending upon your time zone, in which weekday the week begins
	WeeksStartAt() time.Weekday
}
