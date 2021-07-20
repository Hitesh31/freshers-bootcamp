package main

import "fmt"

type Matrix struct {
	numRows, numColumns int
	elements [][]int
}
func (m Matrix) getRows() int{
	return m.numRows
}
func (m Matrix) getColumns() int{
	return m.numColumns
}
func (m *Matrix) setElement(i , j , num int) {
	m.elements[i][j]=num
}
func main() {

	mat := Matrix{numRows: 2,numColumns: 3}
	a := make([][] int, mat.getRows())
	for i:=0; i<mat.getRows();i++ {
		a[i]=make([]int, mat.getColumns())
	}
	fmt.Println(a)
	mat.elements = a
	mat.setElement(1,1,6)
	fmt.Println(mat.elements[1][1])
	fmt.Println("number of rows and Columns are" , mat.getRows(),"and",mat.getColumns(), ".")

	fmt.Println(mat)
}