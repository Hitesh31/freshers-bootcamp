package main

import "fmt"

type CalcSalary interface {
	Salary() int
}
type fullTime struct {
	basicPay int
}
type contracter struct {
	basicPay int
}
type freelancer struct {
	payPerHour int
	numHours int
}
func (f fullTime) Salary() int {
	return 28*f.basicPay
}
func (c contracter) Salary() int {
	return 28*c.basicPay
}
func (f freelancer) Salary() int {
	return f.numHours*f.payPerHour
}
func main() {
	person1 := fullTime{basicPay: 500}
	person2 := contracter{basicPay: 100}
	person3 := freelancer{payPerHour: 10, numHours: 10}
	fmt.Println("fulltime employee monthly salary is", person1.Salary())
	fmt.Println("contracter employee monthly salary is", person2.Salary())
	fmt.Println("freelancer employee daily salary is", person3.Salary())
}