package main

import (
	"fmt"
)

type Simpler interface {
	Get() int
	Set(n int)
}

type Simple struct {
	n int
}

func (s *Simple) Get() int {
	return s.n
}

func (s *Simple) Set(n int) {
	s.n = n
}

type RSimple struct {
	t int
}

func (s *RSimple) Get() int {
	return s.t
}

func (s *RSimple) Set(n int) {
	s.t = n
}

func test(s Simpler) int {
	s.Set(122)
	return s.Get()
}

func fi(s Simpler) string {
	switch s.(type) {
	case *Simple: return "Simple"
	case *RSimple: return "RSimple"
	}
	return "NOT FOUND"
}

func gi(any interface{}) bool {
	switch any.(type) {
	case Simpler : return true
	}
	return false
}

func gi2(any interface{}) int {
	if v, ok := any.(Simpler); ok {
		return v.Get()
	}
	return 0 // default value
}

func main()  {
	var s Simple
	fmt.Println(test(&s)) // &s is required because Get() is defined with a receiver type pointer
	fmt.Println(fi(&s))

	fmt.Printf("s type is %T, gi result: %v\n", s, gi(s))
	fmt.Printf("s type is %T, gi result: %v\n", s, gi(&s))

	fmt.Printf("s type is %T, gi2 value: %v\n", s, gi(s))
	fmt.Printf("s type is %T, gi2 value: %v\n", s, gi(&s))

	/*
	s type is main.Simple, gi result: false
s type is main.Simple, gi result: true
s type is main.Simple, gi2 value: false
s type is main.Simple, gi2 value: true
	 */
}