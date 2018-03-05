package main

import "fmt"

func main()  {
	for i:=1; i<=25; i++ {
		for j:= 1; j < i; j++ {
			fmt.Print("G")
		}
		fmt.Println()
	}

	fmt.Println("-------------------------------------")
	// 截断 一个for
	s := "GGGGGGGGGGGGGGGGGGGGGGGG"
	for i:=1; i<25; i++ {
		fmt.Println(s[0:i])
	}
}
