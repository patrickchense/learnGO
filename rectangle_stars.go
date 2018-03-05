package main

import (
	"fmt"
)

func main()  {
	for i:= 0; i < 10; i++ {
		fmt.Print("*")

			for j:= 0; j<18; j++ {
				if i % 10 == 0 || (i+1) % 10 == 0 {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			}
		fmt.Println("*")
	}
}
