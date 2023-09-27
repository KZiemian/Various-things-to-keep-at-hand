package main

import "fmt"

// type Employee struct {
// 	SalaryPerYear int
// }

// type Freelancer struct {
// 	RatePerHour int
// 	HoursWorked int
// }

// type SalaryCalculator interface {
// 	Salary() int
// }

// type HourlyRateCalculator interface {
// 	HourlyRate() int
// }

// func (e *Employee) Salary() int {
// 	return e.SalaryPerYear
// }

// func (e *Employee) HourlyRate() int {
// 	return e.SalaryPerYear / 12 / 30 / 8
// }

// func (f *Freelancer) HourlyRate() int {
// 	return f.RatePerHour
// }

// func (f *Freelancer) Salary() int {
// 	return f.RatePerHour * f.HoursWorked
// }

// type SalaryCalculations interface {
// 	SalaryCalculator
// 	HourlyRateCalculator
// }

// func CalculateTotalSalaries(scs ...SalaryCalculator) {
// 	total := 0

// 	for _, sc := range scs {
// 		total += sc.Salary()
// 	}

// 	fmt.Printf("Total salaries paid: %d.\n", total)
// }

// func main() {
// 	e1 := &Employee{
// 		SalaryPerYear: 1000,
// 	}

// 	f1 := &Freelancer{
// 		RatePerHour: 200,
// 		HoursWorked: 100,
// 	}

// 	CalculateTotalSalaries(e1, f1)
// }

type Counter interface {
	Count() int
}

type Word string

type Number int

func (w Word) Count() int {
	return len(w)
}

func (n Number) Count() int {
	count := 0

	for n > 0 {
		n = n / 10

		count++
	}

	return count
}

func describe(c Counter) {
	fmt.Printf("Interface type %T value %v.\n", c, c)
}

func main() {
	var w Word = "helloworld"
	var n Number = 123456

	fmt.Printf("w.Count(): %v.\n", w.Count())
	fmt.Printf("n.Count(): %v.\n", n.Count())

	describe(w)
	describe(n)
}

// v, ok := i.(int)
