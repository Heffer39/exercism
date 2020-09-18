// Package clock implements a clock that handles times without dates
package clock

import (
	"fmt"
)

// Clock represents a 24 hour clock that captures the total minutes of a given time
type Clock struct {
	minutes int
}

// New acts as the constructor for the Clock struct
func New(hour, minute int) Clock {
	minute += hour * 60
	minute %= 24 * 60
	if minute < 0 {
		minute += 24 * 60
	}
	return Clock{minute}
}

// Add takes a clock and adds minutes to it
func (c Clock) Add(add int) Clock {
	return New(0, c.minutes+add)
}

// Subtract takes a clock and subtracts minutes from it
func (c Clock) Subtract(subtract int) Clock {
	return New(0, c.minutes-subtract)
}

// String takes a clock and prints out the hours and minutes in "00:00" 24h format
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minutes/60, c.minutes%60)
}
