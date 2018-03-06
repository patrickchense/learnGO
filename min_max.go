package main

import "fmt"

func main()  {
	fmt.Printf("min slice: %d", minSlice([]int{1,2,3,4}))
	fmt.Printf("max slice: %d", maxSlice([]int{1,2,3,4}))
}

func minSlice(mins []int) int {
	var min int = mins[0]
	for _, v := range mins {
		if v < min {
			min = v
		}
	}
	return min
}

func maxSlice(maxs []int) int {
	var max int = maxs[0]
	for _, v := range maxs {
		if v > max {
			max = v
		}
	}
	return max
}
//min slice: 1max slice: 4