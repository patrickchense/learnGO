package main


import (
	"fmt"
	"math"
)

type Point4 struct {
	x, y float64
}

func (p *Point4) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

type NamedPoint struct {
	Point4
	name string
}

func main() {
	n := &NamedPoint{Point4{3, 4}, "Pythagoras"}
	fmt.Println(n.Abs()) // 打印5
}
