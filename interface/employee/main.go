package main

import (
	"fmt"
)

//SalaryCalculator 创建接口
type SalaryCalculator interface {
	CalculateSalary() int
}

//Permanent 结构体
type Permanent struct {
	empID    int
	basicPay int
	pf       int
}

//Contract 结构体
type Contract struct {
	empID    int
	basicPay int
}

//CalculateSalary 实现该方法
func (p Permanent) CalculateSalary() int {
	return p.basicPay + p.pf
}

// CalculateSalary 实现方法
func (c Contract) CalculateSalary() int {
	return c.basicPay
}

func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total Expense per Month $%d\n", expense)
}

func main() {
	pemp1 := Permanent{1, 5000, 20}
	pemp2 := Permanent{2, 6000, 30}
	cemp1 := Contract{3, 3000}
	employee := []SalaryCalculator{pemp1, pemp2, cemp1}
	totalExpense(employee)
}
