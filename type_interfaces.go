package main

import (
	"fmt"
	"math"
)

type Square1 struct {
	side float32
}

type Circle1 struct {
	radius float32
}

type Shaper1 interface {
	Area() float32
}

func main() {
	var areaIntf Shaper1
	sq1 := new(Square1)
	sq1.side = 5

	areaIntf = sq1
	// Is Square the type of areaIntf?
	if t, ok := areaIntf.(*Square1); ok {
		fmt.Printf("The type of areaIntf is: %T\n", t)
	}
	if u, ok := areaIntf.(*Circle1); ok {
		fmt.Printf("The type of areaIntf is: %T\n", u)
	} else {
		fmt.Println("areaIntf does not contain a variable of type Circle")
	}
}

func (sq *Square1) Area() float32 {
	return sq.side * sq.side
}

func (ci *Circle1) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}

