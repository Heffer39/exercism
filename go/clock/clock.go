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

// New acts as the constructor for the Clock type
func New(hour, minute int) Clock {
	c := new(Clock)
	c.hour, c.minute = norm(hour, minute, 60)
	_, c.hour = norm(1, c.hour, 24)
	return *c
}

// norm returns nhi, nlo such that
//	hi * base + lo == nhi * base + nlo
//	0 <= nlo < base
// Copied private method norm from Package time
func norm(hi, lo, base int) (nhi, nlo int) {
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

// Add takes a clock and adds minutes to it
func (c Clock) Add(add int) Clock {
	c.hour, c.minute = norm(c.hour, c.minute+add, 60)
	_, c.hour = norm(0, c.hour, 24)
	return c
}

// Subtract takes a clock and subtracts minutes from it
func (c Clock) Subtract(subtract int) Clock {
	c.hour, c.minute = norm(c.hour, c.minute-subtract, 60)
	_, c.hour = norm(0, c.hour, 24)
	return c
}

// String takes a clock and prints out the hours and minutes in "00:00" 24h format
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
