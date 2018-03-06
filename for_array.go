package main

import "fmt"

func main()  {
	var a [16]int
	for i := range a {
		a[i] = i
		fmt.Printf("a[%d] = %d", i, a[i])
		fmt.Println()
	}
}
/**
a[0] = 0
a[1] = 1
a[2] = 2
a[3] = 3
a[4] = 4
a[5] = 5
a[6] = 6
a[7] = 7
a[8] = 8
a[9] = 9
a[10] = 10
a[11] = 11
a[12] = 12
a[13] = 13
a[14] = 14
a[15] = 15
 */
