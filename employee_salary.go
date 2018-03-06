package main

import "fmt"

type employee struct {
	salary int
}

func (e *employee) giveRaise(factor float64) {
	e.salary = int(float64(e.salary) * factor)
}

func main()  {
	e := new(employee)
	e.salary = 100
	e.giveRaise(1.1)
	fmt.Printf("e new salary:%d", e.salary) //e new salary:110
}