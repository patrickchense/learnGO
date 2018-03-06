package main

import "fmt"

func main()  {
	week := map[string]int {"Monday":1, "Tuesday":2, "Wesday":3, "Thesday":4, "Friday":5, "Satuday":6, "Sunday":7}
	if _, isPresent := week["Tuesday"]; isPresent {
		fmt.Println("Tuesday exists")
	} else {
		fmt.Println("Tuesday not exists")
	}

	if _, isPresent := week["Hollyday"]; isPresent {
		fmt.Println("Hollyday exists")
	} else {
		fmt.Println("Hollyday not exists")
	}

}
/*
Tuesday exists
Hollyday not exists
 */