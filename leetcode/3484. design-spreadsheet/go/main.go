package leetcode3484

import (
	"strconv"
)

type Spreadsheet struct {
	rows map[string][]int
}

func Constructor(rows int) Spreadsheet {
	spreadsheetRows := make(map[string][]int)

	for c := byte('A'); c <= byte('Z'); c++ {
		spreadsheetRows[string(c)] = make([]int, rows)
	}

	return Spreadsheet{
		rows: spreadsheetRows,
	}
}

func (s *Spreadsheet) SetCell(cell string, value int) {
	col := s.parseCol(cell)
	row := s.parseRow(cell)

	s.setValue(string(col), row, value)
}

func (s *Spreadsheet) ResetCell(cell string) {
	col := s.parseCol(cell)
	row := s.parseRow(cell)

	s.setValue(string(col), row, 0)
}

func (s *Spreadsheet) GetValue(formula string) int {
	if formula[0] != '=' {
		return 0
	}

	sum := 0
	leftOperand := ""
	rightOperand := ""
	isFirstOperand := true
	for i := 1; i < len(formula); i++ {
		r := formula[i]
		if r == '+' {
			isFirstOperand = false
			continue
		}

		if isFirstOperand {
			leftOperand += string(r)
		} else {
			rightOperand += string(r)
		}
	}

	left := s.getValueByFormulaPart(leftOperand)
	right := s.getValueByFormulaPart(rightOperand)

	sum = left + right

	return sum
}

func (s *Spreadsheet) getValueByFormulaPart(part string) int {
	if part[0] >= 'A' && part[0] <= 'Z' {
		col := s.parseCol(part)
		row := s.parseRow(part)

		return s.getValue(string(col), row)
	}

	value, _ := strconv.Atoi(part)

	return value
}

func (s *Spreadsheet) parseCol(col string) byte {
	return col[0]
}

func (s *Spreadsheet) parseRow(row string) int {
	n, err := strconv.Atoi(row[1:])
	if err != nil {
		return 0
	}

	return n
}

func (s *Spreadsheet) getValue(col string, row int) int {
	rows := s.rows[col]
	if row <= 0 || row > len(rows) {
		return 0
	}

	return s.rows[col][row-1]
}

func (s *Spreadsheet) setValue(col string, row int, value int) {
	rows := s.rows[col]
	if row <= 0 || row > len(rows) {
		return
	}

	s.rows[col][row-1] = value
}

/**
 * Your Spreadsheet object will be instantiated and called as such:
 * obj := Constructor(rows);
 * obj.SetCell(cell,value);
 * obj.ResetCell(cell);
 * param_3 := obj.GetValue(formula);
 */
