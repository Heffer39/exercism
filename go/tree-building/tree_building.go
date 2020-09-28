// Package tree is responsible for converting a collection of records into a tree-based structure of Nodes
package tree

import (
	"errors"
	"sort"
)

// Record is a struct containing int fields ID and Parent
type Record struct {
	ID, Parent int
}

// Node is a struct containing int field ID and []*Node field Children
type Node struct {
	ID       int
	Children []*Node
}

// Build takes in a slice of records, which it uses to create a tree-based hierarchy of Nodes
// It returns a reference to the root node from the constructed tree
func Build(records []Record) (*Node, error) {
	nodeMap := make(map[int]*Node, len(records))

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	for i, r := range records {
		if r.ID != i || r.Parent > r.ID || r.ID > 0 && r.Parent == r.ID {
			return nil, errors.New("invalid record")
		}
		nodeMap[r.ID] = &Node{ID: r.ID}
		if i != 0 {
			parentNode := nodeMap[r.Parent]
			parentNode.Children = append(parentNode.Children, nodeMap[r.ID])
		}
	}

	return nodeMap[0], nil
}
