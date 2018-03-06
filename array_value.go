package main

import "fmt"

func main()  {
	var a [3]int = [3]int{1,2,3}
	var b = a
	for i := range b {
		b[i] = b[i] * 2
	}
	fmt.Println(a)
	fmt.Println(b)
}
/*
[1 2 3]
[2 4 6]
 */