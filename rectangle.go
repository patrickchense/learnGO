package main

import "fmt"

type Rectangle struct {
	X, Y int
}

func Area(r *Rectangle) int {
	return r.X * r.Y
}

func Primeter(r *Rectangle) int {
	return 2*r.Y + 2*r.X
}

func main()  {
	r1 := &Rectangle{100, 1}
	fmt.Printf("r1 Area : %d\n", Area(r1))
	fmt.Printf("r1 Primerter : %d \n", Primeter(r1))
}
/*
r1 Area : 100
r1 Primerter : 202
 */