package main

import "fmt"

type Engine interface {
	Start()
	Stop()
}



type Car struct {
	Engine
	wheelCount int
}

type Mercedes struct {
	Car
}

func (c *Car) numberOfWheels() int {
	return c.wheelCount
}

func (m *Mercedes) sayHiToMerkel() {
	println("say hi to merkel")
}

func (c *Car) Start() {
	fmt.Println("Car is started")
}

func (c *Car) Stop() {
	fmt.Println("Car is stopped")
}

func (c *Car) GoToWorkIn() {
	// get in car
	c.Start()
	// drive to work
	c.Stop()
	// get out of car
}

func main()  {
	m := Mercedes{Car{nil, 4}}
	fmt.Println("A Mercedes has this many wheels: ", m.numberOfWheels())
	m.GoToWorkIn()
	m.sayHiToMerkel()
}


