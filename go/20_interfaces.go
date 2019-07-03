package main

import "fmt"

type salary interface{
	calculateSalary() float64
	calculateTax() float64
}

type employee struct{
	perDayIncome float64
}

type manager struct{
	perDayIncome float64
	incentive int
}

// func (e employee) calculateSalary() float64{		error
func (e employee) calculateSalary() float64{
		return e.perDayIncome * 365
}

// func (e *employee) calculateTax() float64{		error
func (e employee) calculateTax() float64{
	return e.calculateSalary() * 30 /100
}

// func (m manager) calculateSalary() float64{		error
func (m manager) calculateSalary() float64{
	return (m.perDayIncome* 365) + (float64(m.incentive) * 12)
}

// func (m manager) calculateTax() float64{			error
func (m manager) calculateTax() float64{
	return m.calculateSalary() * 35 / 100
}

func calculate(s salary){
	fmt.Println("object", s)
	fmt.Println("calculated salary", s.calculateSalary())
	fmt.Println("calculated tax", s.calculateTax())
}

func main(){
	e := employee{perDayIncome: 1000}
	m := manager{perDayIncome: 2000, incentive: 5000}

	calculate(e)
	fmt.Println("\n\n")
	calculate(m)
}
