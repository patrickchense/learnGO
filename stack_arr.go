package main

import (
	"strconv"
	"fmt"
)

const LIMIT1 = 4

type Stack1 [LIMIT1]int

func (st *Stack1) Push(n int) {
	for ix, v := range st {
		if v == 0 {
			st[ix] = n
			break
		}
	}
}

// take value from first position which contains !=0, starting from top
func (st *Stack1) Pop() int {
	v := 0
	for ix := len(st) - 1; ix >= 0; ix-- {
		if v = st[ix]; v != 0 {
			st[ix] = 0
			return v
		}
	}
	return 0
}

func (st Stack1) String() string {
	str := ""
	for ix, v := range st {
		str += "[" + strconv.Itoa(ix) + ":" + strconv.Itoa(v) + "] "
	}
	return str
}

func main() {
	st1 := new(Stack1)
	fmt.Printf("%v\n", st1)
	st1.Push(3)
	fmt.Printf("%v\n", st1)
	st1.Push(7)
	fmt.Printf("%v\n", st1)
	st1.Push(10)
	fmt.Printf("%v\n", st1)
	st1.Push(99)
	fmt.Printf("%v\n", st1)
	p := st1.Pop()
	fmt.Printf("Popped %d\n", p)
	fmt.Printf("%v\n", st1)
	p = st1.Pop()
	fmt.Printf("Popped %d\n", p)
	fmt.Printf("%v\n", st1)
	p = st1.Pop()
	fmt.Printf("Popped %d\n", p)
	fmt.Printf("%v\n", st1)
	p = st1.Pop()
	fmt.Printf("Popped %d\n", p)
	fmt.Printf("%v\n", st1)
}