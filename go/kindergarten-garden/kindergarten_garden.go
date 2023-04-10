// Package kindergarten contains a garden flower project for a kindergarten class
package kindergarten

import (
	"bufio"
	"errors"
	"fmt"
	"sort"
	"strings"
)

// Garden represents the diagram of plants in a garden and the children they belong to
type Garden struct {
	children       []string
	childrenPlants map[string][]string
}

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`

// grass, clover, radishes, and violets
var plantNames = map[string]string{
	"G": "grass",
	"C": "clover",
	"R": "radishes",
	"V": "violets",
}

// NewGarden generates a new Garden of flowers on a windowsill, with each flower belonging to a child
func NewGarden(diagram string, children []string) (*Garden, error) {
	if strings.Count(diagram, "\n") != 2 {
		return nil, errors.New("wrong diagram format")
	}
	if len(children) == 0 {
		return nil, errors.New("no children")
	}

	garden := Garden{
		childrenPlants: map[string][]string{},
	}

	allKeys := make(map[string]bool)
	var listOfChildren []string
	for _, item := range children {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			listOfChildren = append(listOfChildren, item)
		} else {
			return nil, errors.New("mismatched rows")
		}
	}
	garden.children = listOfChildren
	sort.Strings(garden.children)

	var lines []string
	scanner := bufio.NewScanner(strings.NewReader(diagram))
	scanner.Scan()
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var letterCount, childCount, rowLength int
	for i, line := range lines {
		if i > 0 {
			if rowLength != len(line) {
				return nil, errors.New("mismatched rows")
			}
		}
		rowLength = len(line)
		if rowLength/2 > len(children) {
			return nil, errors.New("not enough children")
		}
		if rowLength%2 != 0 {
			return nil, errors.New("odd cups")
		}
		fmt.Println(line)
		line = strings.Replace(line, "\n", "", 1)
		for _, character := range line {
			if letterCount == 2 {
				childCount++
				letterCount = 0
			}
			plantKeys := keys(plantNames)
			if contains(plantKeys, fmt.Sprintf("%c", character)) == false {
				return nil, errors.New("invalid plant code")
			}
			garden.childrenPlants[garden.children[childCount]] = append(garden.childrenPlants[garden.children[childCount]],
				plantNames[fmt.Sprintf("%c", character)])
			letterCount++
		}
		childCount = 0
		letterCount = 0
	}

	return &garden, nil
}

// Plants returns the list of plants that belongs to the given child
func (g *Garden) Plants(child string) ([]string, bool) {
	if g.childrenPlants[child] != nil {
		return g.childrenPlants[child], true
	}
	return nil, false
}

func keys(stringMap map[string]string) (keys []string) {
	for i := range stringMap {
		keys = append(keys, i)
	}
	return keys
}

func contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
