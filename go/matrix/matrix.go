// Package matrix implements matrix operations
package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

// Matrix represents a two dimension slice of ints
type Matrix [][]int

// New returns a new two dimensional int matrix from a given string
func New(s string) (Matrix, error) {
	rows := strings.Split(s, "\n")
	matrix := make([][]int, len(rows))
	lenCols := len(strings.Split(strings.Trim(rows[0], " "), " "))
	for i, row := range rows {
		nums := strings.Split(strings.Trim(row, " "), " ")
		if len(nums) != lenCols {
			return nil, fmt.Errorf("invalid matrix dimensions")
		}
		matrix[i] = make([]int, len(nums))
		for j, v := range nums {
			num, err := strconv.Atoi(v)
			if err != nil {
				return nil, fmt.Errorf("invalid input %s", v)
			}
			matrix[i][j] = num
		}
	}
	return matrix, nil
}

// Cols returns a list of columns from the original matrix, reading each column from top to bottom
// while moving from left to right across the rows
func (m *Matrix) Cols() [][]int {
	matrix := make([][]int, len((*m)[0]))
	for i, row := range *m {
		for j, v := range row {
			if i == 0 {
				matrix[j] = make([]int, len(*m))
			}
			matrix[j][i] = v
		}
	}
	return matrix
}

// Rows returns a list of rows in the original matrix, reading each row from left to right
// while moving top to bottom across the rows
func (m *Matrix) Rows() [][]int {
	matrix := make([][]int, len(*m))
	for i, row := range *m {
		matrix[i] = make([]int, len(row))
		for j, v := range row {
			matrix[i][j] = v
		}
	}
	return matrix
}

// Set places the given value in the specified [row][column] coordinates
func (m *Matrix) Set(row int, column int, value int) bool {
	if row < 0 || column < 0 || row > len(*m)-1 || column > len((*m)[row])-1 {
		fmt.Println("invalid input")
		return false
	}
	(*m)[row][column] = value
	return true
}
