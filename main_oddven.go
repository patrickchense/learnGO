package main

import (
	"./evens"
	"fmt"
)

func main() {
	for i := 0; i <= 100; i++ {
		fmt.Printf("Is the integer %d even? %v\n", i, evens.Even(i))
	}
}