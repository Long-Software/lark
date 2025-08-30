package math

import "fmt"

type Matrix []Vector

func NewMatrix(row, col int) Matrix {
	m := make([]Vector, row)
	for i := 0; i < row; i++ {
		m[i] = make([]Scalar, col)
	}
	return m
}

func (m *Matrix) InterChange(row_n, row_m int) error {
	row, _ := m.Size()
	if row_n >= row || row_n < 0 || row_m >= row || row_m < 0 {
		return OutOfBound
	}
	(*m)[row_n], (*m)[row_m] = (*m)[row_m], (*m)[row_n]
	return nil
}
func (m *Matrix) RowMul(i int, k Scalar) error {
	row, _ := m.Size()
	if i >= row || i < 0 {
		return OutOfBound
	}
	(*m)[i] = (*m)[i].Times(k)
	return nil
}

func (m *Matrix) AddRow(i int, r Vector) error {
	row, col := m.Size()
	if len(r) != col {
		return InCompatibleLengthError
	}
	if i >= row || i < 0 {
		return OutOfBound
	}
	vector, err := vec.Add((*m)[i], r)
	if err != nil {
		return err
	}
	(*m)[i] = vector
	return nil
}

func (m *Matrix) Size() (int, int) {
	return len(*m), len((*m)[0])
}

func (m *Matrix) String() string {
	x, y := m.Size()
	var str string
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			str += fmt.Sprintf("%v ", (*m)[i][j])
		}
		str += "\n"
	}
	return str
}
