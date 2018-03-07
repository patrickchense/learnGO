package main

import (
	"math"
	"fmt"
)

type Point5 struct {
	X, Y float64
}

func (p *Point5) Scale(s float64) {
	p.X *= s
	p.Y *= s
}

func (p *Point5) Abs() float64 {
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}

type Point6 struct {
	X, Y, Z float64
}

func (p *Point6) Abs() float64 {
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y + p.Z*p.Z))
}

type Polar1 struct {
	R, T float64
}

func (p Polar1) Abs() float64 { return p.R }

func main() {
	p1 := new(Point5)
	p1.X = 3
	p1.Y = 4
	fmt.Printf("The length of the vector p1 is: %f\n", p1.Abs())

	p2 := &Point5{4, 5}
	fmt.Printf("The length of the vector p2 is: %f\n", p2.Abs())

	p1.Scale(5)
	fmt.Printf("The length of the vector p1 after scaling is: %f\n", p1.Abs())
	fmt.Printf("Point p1 after scaling has the following coordinates: X %f - Y %f", p1.X, p1.Y)
}
