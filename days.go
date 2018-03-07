package main

import "fmt"

type Day int

const (
	MO = iota
	TU
	WE
	TH
	FR
	ST
	SU
)

func (d Day) String() string {
	switch d % 7 {
	case MO: return "MO"
	case TU: return "TU"
	case WE: return "WE"
	case TH: return "TH"
	case FR: return "FR"
	case ST: return "ST"
	case SU: return "SU"
	}
	return "NOT FOUND"
}

func main()  {
	var th Day = 3
	fmt.Printf("The 3rd day is: %s\n", th)
	// If index > 6: panic: runtime error: index out of range
	// but use the enumerated type to work with valid values:
	var day = SU
	fmt.Println(day) // prints Sunday
	fmt.Println(0, MO, 1, TU)
}
/*
The 3rd day is: TH
6
0 0 1 1
 */