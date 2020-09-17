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

// New acts as the constructor for the Clock type and
// returns hours and minutes in a 24 hour, 60 minute clock format
func New(hour, minute int) Clock {
	c := new(Clock)
	base := 60

	if minute < 0 {
		n := (-minute-1)/base + 1
		hour -= n
		minute += n * base
	}
	if minute >= base {
		n := minute / base
		hour += n
		minute -= n * base
	}

	c.hour, c.minute = hour, minute
	c.hour = (24 + c.hour%24) % 24
	return *c
}

// Add takes a clock and adds minutes to it
func (c Clock) Add(add int) Clock {
	return New(c.hour, c.minute+add)
}

// Subtract takes a clock and subtracts minutes from it
func (c Clock) Subtract(subtract int) Clock {
	return New(c.hour, c.minute-subtract)
}

// String takes a clock and prints out the hours and minutes in "00:00" 24h format
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
