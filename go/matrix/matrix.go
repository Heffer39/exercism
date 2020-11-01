package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

type Matrix [][]int

func New(s string) (Matrix, error) {
	//fmt.Printf("Problem:\n%s\n", s)
	rows := strings.Split(s, "\n")
	//_, _ = fmt.Printf("rows: %s\n", rows)
	matrix := make([][]int, len(rows))
	for i, row := range rows {
		row = strings.Trim(row, " ")
		nums := strings.Split(row, " ")
		//_, _ = fmt.Printf("row: %s\n", row)
		//_, _ = fmt.Println(nums)
		matrix[i] = make([]int, len(nums))
		for j, v := range nums {
			if v == " " {
				continue
			}
			num, err := strconv.Atoi(v)
			if err != nil {
				return nil, fmt.Errorf("invalid input %s", v)
			}
			matrix[i][j] = num
			/*if !unicode.IsNumber(v) && !unicode.IsSpace(v) {
				return nil, fmt.Errorf("invalid input")
			}
			matrix[i][j] = int(v - '0')*/
		}
	}
	//fmt.Printf("Matrix:\n%v\n", matrix)

	var rowSize int
	for _, row := range matrix {
		if rowSize == 0 {
			rowSize = len(row)
		}
		if rowSize != len(row) {
			return nil, fmt.Errorf("invalid matrix dimensions")
		}
		//fmt.Printf("row size: %v\n", len(row))
	}
	return matrix, nil
}

func (m *Matrix) Cols() [][]int {
	matrix := make([][]int, len((*m)[0]))
	for i, row := range *m {
		for j, v := range row {
			if i == 0 {
				matrix[j] = make([]int, len(*m))
			}
			//fmt.Printf("val %v\n", v)
			matrix[j][i] = v
		}
	}
	//fmt.Printf("Col matrix:\n%v\n", matrix)
	return matrix
}

func (m *Matrix) Rows() [][]int {
	matrix := make([][]int, len(*m))
	//fmt.Printf("row: matrix length %v\n", len(matrix))
	for i, row := range *m {
		//fmt.Printf("row: matrix row length %v\n", len(row))
		matrix[i] = make([]int, len(row))
		for j, v := range row {
			matrix[i][j] = v
		}
	}
	return matrix
}

func (m *Matrix) Set(row int, column int, value int) bool {
	//fmt.Printf("Set matrix:\n%v\n", *m)
	if row < 0 || column < 0 || row > len(*m)-1 || column > len((*m)[row])-1 {
		fmt.Println("bad input")
		return false
	}
	(*m)[row][column] = value
	//fmt.Printf("Set matrix update:\n%v\n", *m)
	//fmt.Printf("row: %v, col: %v, val: %v\n", row, column, value)
	return true
}
