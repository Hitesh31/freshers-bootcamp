package main

import (
	"encoding/json"
	"fmt"
)

type Matrix struct {
	Rows int `json:"Rows"`
	Columns int `json:"Columns"`
	Elements [][]int `json:"Elements"`
}
func (m Matrix) getRows() int{
	return m.Rows
}
func (m Matrix) getColumns() int{
	return m.Columns
}
func (m *Matrix) setElement(i , j , num int) {
	m.Elements[i][j]=num
}
func addMatrix(m Matrix, n Matrix) Matrix {
	var addMat Matrix
	addMat = n
	for i:=0;i<m.Rows;i++ {
		for j:=0;j<m.Columns;j++ {
			addMat.Elements[i][j]=m.Elements[i][j] + addMat.Elements[i][j]
		}
	}
	return addMat
}

func (m *Matrix) toJson(){
	k, err :=json.MarshalIndent(m,""," ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(k))
}

func main() {

	mat := Matrix{Rows: 2,Columns: 3}
	a := make([][] int, mat.getRows())
	for i:=0; i<mat.getRows();i++ {
		a[i]=make([]int, mat.getColumns())
	}
	mat.Elements = a
	mat.setElement(1,1,6)

	var mat2 Matrix
	b := make([][] int, mat.getRows())
	for i:=0; i<mat.getRows();i++ {
		b[i]=make([]int, mat.getColumns())
	}
	mat2.Elements = b
	mat2.setElement(1,0,7)

	j := addMatrix(mat,mat2)
	fmt.Println(j)
	j.toJson()
}