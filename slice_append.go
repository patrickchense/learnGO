package main

import "fmt"

func main()  {
	// init
	s := make([]byte, 3, 4)
	for i:=0; i<len(s); i++ {
		s[i] = 'a'
	}
	fmt.Println(s)
	s = Append(s, []byte{'b'})
	fmt.Println(s)
	s = Append(s, []byte{'c','d'})
	fmt.Println(s)
}

func Append(s []byte, data []byte) []byte {
	if cap(s) - len(s) >= len(data) {
		s = append(s, data ...)
	} else {
		fmt.Printf("enlarge slice %d -> %d", cap(s), cap(s) + len(data))
		ns := make([]byte, len(s), cap(s) + len(data) - (cap(s) - len(s)))
		copy(ns, s)
		s = append(ns, data ...)
	}
	return s
}
/*
[97 97 97]
[97 97 97 98]
enlarge slice 4 -> 6[97 97 97 98 99 100]
 */