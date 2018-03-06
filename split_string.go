package main

import "fmt"

func main()  {
	s := "akadlfkdjfklsdf"
	s1, s2 := split(s, 6)
	fmt.Printf("s1=%s, s2=%s", s1, s2)
}

func split(s string, idx int) (s1 string , s2 string)  {
	cs := []byte(s)
	s1 = string(cs[0:idx])
	s2 = string(cs[idx:])
	return
}
//s1=akadlf, s2=kdjfklsdf