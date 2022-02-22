package utils

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func RemoveRow(i int, matrix *mat.Dense) error {
	r, c := matrix.Dims()
	if i < 0 || i >= r {
		return fmt.Errorf("deleted row index must be >=0 and less than matrix row number")
	}
	m1 := matrix.Slice(0, i, 0, c)
	m2 := matrix.Slice(i+1, r, 0, c)
	matrix.Reset()
	matrix.Stack(m1, m2)
	return nil
}

func RemoveCol(i int, matrix *mat.Dense) error {
	r, c := matrix.Dims()
	if i < 0 || i >= c {
		return fmt.Errorf("deleted col index must be >=0 and less than matrix col number")
	}
	m1 := matrix.Slice(0, r, 0, i)
	m2 := matrix.Slice(0, r, i+1, c)
	matrix.Reset()
	matrix.Augment(m1, m2)
	return nil
}

func RemoveRowCol(i, j int, matrix *mat.Dense) error {
	r, c := matrix.Dims()
	if i < 0 || i >= r {
		return fmt.Errorf("deleted row index must be >=0 and less than matrix row number")
	}
	if i < 0 || i >= c {
		return fmt.Errorf("deleted col index must be >=0 and less than matrix col number")
	}
	m1 := matrix.Slice(0, i, 0, j)
	m2 := matrix.Slice(i+1, r, 0, j)
	m3 := matrix.Slice(0, i, j+1, c)
	m4 := matrix.Slice(i+1, r, j+1, c)
	upper := &mat.Dense{}
	upper.Augment(m1, m3)
	lower := &mat.Dense{}
	lower.Augment(m2, m4)
	matrix.Reset()
	matrix.Stack(upper, lower)
	return nil
}
