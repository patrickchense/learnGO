package main

import "fmt"

func main() {
	result := 0
	idx := 0
	for i := 0; i <= 10; i++ {
		result, idx = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d, idx: %d \n", i, result, idx)
	}
}

func fibonacci(n int) (res int, idx int) {
	if n <= 1 {
		res = 1
	} else {
		res1, _ := fibonacci(n-1)
		res2, _ := fibonacci(n-2)
		res = res1 + res2
	}
	idx = n
	return
}
/*
fibonacci(0) is: 1, idx: 0
fibonacci(1) is: 1, idx: 1
fibonacci(2) is: 2, idx: 2
fibonacci(3) is: 3, idx: 3
fibonacci(4) is: 5, idx: 4
fibonacci(5) is: 8, idx: 5
fibonacci(6) is: 13, idx: 6
fibonacci(7) is: 21, idx: 7
fibonacci(8) is: 34, idx: 8
fibonacci(9) is: 55, idx: 9
fibonacci(10) is: 89, idx: 10
 */