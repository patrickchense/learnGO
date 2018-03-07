package main

import "fmt"

func main() {
	mf:= func(any interface{}) interface{} {
		switch any.(type) {
		case int : return any.(int) * 10
		case string : return any.(string) + any.(string)
		}
		return any
	}
	isl := []interface{}{0, 1, 2, 3, 4, 5}
	res1 := mapFunc1(mf, isl)
	for _, v := range res1 {
		fmt.Println(v)
	}
	println()

	ssl := []interface{}{"0", "1", "2", "3", "4", "5"}
	res2 := mapFunc1(mf, ssl)
	for _, v := range res2 {
		fmt.Println(v)
	}
}

func mapFunc1(mf func(interface{}) interface{}, list []interface{}) []interface{} {
	result := make([]interface{}, len(list))

	for ix, v := range list {
		result[ix] = mf(v)
	}

	// Equivalent:
	/*
		for ix := 0; ix<len(list); ix++ {
			result[ix] = mf(list[ix])
		}
	*/
	return result
}
/*
0

10
20
30
40
50
00
11
22
33
44
55
 */