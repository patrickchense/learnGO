package main

import "fmt"

func main()  {
	for i := 0; i < 15; i++ {
		fmt.Println(i)
	}

	i := 0
	Start:
		fmt.Println(i)
		if i < 15 {
			i++
			goto Start
		}

}
