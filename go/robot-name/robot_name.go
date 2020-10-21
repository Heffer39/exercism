package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Robot contains the name for a robot
type Robot struct {
	name string
}

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

var usedNames = map[string]bool{}

// Name generates a name for the given robot while verifying that
// the randomly generated name hasn't been duplicated and all names have not been assigned
func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}
	if len(usedNames) == 26*26*10*10*10 {
		return "", errors.New("all possible names have been used")
	}
	r.name = generateName()
	for usedNames[r.name] {
		r.name = generateName()
	}
	usedNames[r.name] = true

	return r.name, nil
}

// generateName populates a Robot's name as a 5 character string
// The first two characters are uppercase letters
// The last three characters are numbers
func generateName() string {
	r1 := seededRand.Intn(26) + 'A'
	r2 := seededRand.Intn(26) + 'A'
	num := seededRand.Intn(1000)
	return fmt.Sprintf("%c%c%03d", r1, r2, num)
}

// Reset simply sets the Robot's name to a blank string
func (r *Robot) Reset() {
	r.name = ""
}
