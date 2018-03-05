package main

import (
	"unicode/utf8"
	"fmt"
)

func main()  {
	str := "asSASA ddd dsjkdsjs dk"
	str1 := "asSASA ddd dsjkdsjsこん dk"
	fmt.Printf("str bytes len : %d\n", len(str))
	fmt.Printf("str1 bytes len : %d\n", len(str1))
	fmt.Printf("str character len : %d\n",utf8.RuneCountInString(str))
	fmt.Printf("str1 character len : %d\n",utf8.RuneCountInString(str1))
}
/*
str bytes len : 22
str1 bytes len : 28
str character len : 22
str1 character len : 24  //日文时3个字节 = 一个字符
 */