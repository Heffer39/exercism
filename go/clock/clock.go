// Package clock implements a clock that handles times without dates
package clock

import (
	"fmt"
)

// Clock represents a 24 hour clock that captures the hour and minute of a given time
type Clock struct {
	hour   int
	minute int
}

// norm returns nhi, nlo such that
//	hi * base + lo == nhi * base + nlo
//	0 <= nlo < base
// Copied private method norm from Package time
func norm(hi, lo, base int) (int, int) {
	if lo < 0 {
		n := (-lo-1)/base + 1
		hi -= n
		lo += n * base
	}
	if lo >= base {
		n := lo / base
		hi += n
		lo -= n * base
	}
	return hi, lo
}

// normHoursAndMinutes calls the norm function and returns hours and minutes
// in a 24 hour, 60 minute clock format
func normHoursAndMinutes(hour, minute int) (int, int) {
	hour, minute = norm(hour, minute, 60)
	hour = (24 + hour%24) % 24
	return hour, minute
}

// New acts as the constructor for the Clock type
func New(hour, minute int) Clock {
	c := new(Clock)
	c.hour, c.minute = normHoursAndMinutes(hour, minute)
	return *c
}

// Add takes a clock and adds minutes to it
func (c Clock) Add(add int) Clock {
	c.hour, c.minute = normHoursAndMinutes(c.hour, c.minute+add)
	return c
}

// Subtract takes a clock and subtracts minutes from it
func (c Clock) Subtract(subtract int) Clock {
	c.hour, c.minute = normHoursAndMinutes(c.hour, c.minute-subtract)
	return c
}

// String takes a clock and prints out the hours and minutes in "00:00" 24h format
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
