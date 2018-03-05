package main

import "fmt"

func main()  {
	PrintlnVars("a", "c", "df", "eg")
}

func PrintlnVars(a ... string)  {
	for _, val := range a {
		fmt.Println(val)
	}
}
/**
a
c
df
eg
 */