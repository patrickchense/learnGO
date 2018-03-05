package main

import "fmt"

func main()  {
	fmt.Println(mult_returnval(1, 2))
	fmt.Println(mult_returnval2(1, 2))
}

func mult_returnval(a int, b int) (int, int, int) {
	return a + b, a*b, a-b
}

func mult_returnval2(a int, b int) (sum int, mul int, sub int) {
	sum = a + b
	mul = a * b
	sub = a - b
	// sum, mul, sub = a+b, a*b, a-b
	return
}
/**
3 2 -1
3 2 -1
 */