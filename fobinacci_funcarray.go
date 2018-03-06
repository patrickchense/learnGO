package main

import "fmt"

var term = 15

func main() {
	result := fibarray(term)
	for ix, fib := range result {
		fmt.Printf("The %d-th Fibonacci number is: %d\n", ix, fib)
	}
}

func fibarray(term int) []int {
	farr := make([]int, term)
	farr[0], farr[1] = 1, 1

	for i := 2; i < term; i++ {
		farr[i] = farr[i-1] + farr[i-2]
	}
	return farr
}
/*
The 0-th Fibonacci number is: 1
The 1-th Fibonacci number is: 1
The 2-th Fibonacci number is: 2
The 3-th Fibonacci number is: 3
The 4-th Fibonacci number is: 5
The 5-th Fibonacci number is: 8
The 6-th Fibonacci number is: 13
The 7-th Fibonacci number is: 21
The 8-th Fibonacci number is: 34
The 9-th Fibonacci number is: 55
The 10-th Fibonacci number is: 89
The 11-th Fibonacci number is: 144
The 12-th Fibonacci number is: 233
The 13-th Fibonacci number is: 377
The 14-th Fibonacci number is: 610
 */
