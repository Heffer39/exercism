package robotname

import (
	"errors"
	"math/rand"
	"time"
)

// Robot contains the name for a robot
type Robot struct {
	name string
}

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

var usedNames = map[string]bool{}

// Name generates a name for the given robot while verifying that
// the randomly generated name hasn't been duplicated and all names have not been assigned
func (r *Robot) Name() (string, error) {
	if len(usedNames) == 26*26*10*10*10 {
		return "", errors.New("all possible names have been used")
	}

	if r.name == "" {
		r.name = generateName()
		for usedNames[r.name] {
			r.name = generateName()
		}
		usedNames[r.name] = true
	}

	return r.name, nil
}

// generateName populates a Robot's name as a 5 character string
// The first two characters are uppercase letters
// The last three characters are numbers
func generateName() string {
	b := make([]byte, 5)
	for i := 0; i < 5; i++ {
		if i < 2 {
			b[i] = byte(seededRand.Intn(26) + 'A')
		} else {
			b[i] = byte(seededRand.Intn(10) + '0')
		}
	}
	return string(b)
}

// Reset simply sets the Robot's name to a blank string
func (r *Robot) Reset() {
	r.name = ""
}
