package main

import "fmt"

func main()  {
	s := []int{1,2,3,4,5}
	fmt.Printf(" s:%v , after mulTen, s:%v", s, mapFunc(mulTen, s))
}

func mapFunc(fn func(n int) int, ns []int) []int  {
	nns := make([]int, len(ns))
	for i, v := range ns {
		nns[i] = fn(v)
	}
	return nns
}

func mulTen(n int) (nn int) {
	nn = n * 10
	return
}
// s:[1 2 3 4 5] , after mulTen, s:[10 20 30 40 50]