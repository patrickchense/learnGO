package main

import "fmt"
func main()  {
	var mon int = 9
	season := Season(mon)
	fmt.Printf("mon %d is %s \n", mon, season)
	mon = 11
	fmt.Printf("mon %d is %s \n", mon, season)

}

func Season(mon int) string {
	switch mon {
	case 1,2,3 : return "spring"
	case 4,5,6 : return "summer"
	case 7,8,9 : return "fall"
	case 10,11,12 : return "winter"
	default:
		return "error"
	}
}
/*
mon 9 is fall
mon 11 is fall
 */