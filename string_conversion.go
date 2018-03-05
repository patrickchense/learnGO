package main

import (
	"fmt"
	"strconv"
)

/*
strconv.Itoa(i int) string 返回数字 i 所表示的字符串类型的十进制数
strconv.FormatFloat(f float64, fmt byte, prec int, bitSize int) string 将 64 位浮点型的数字转换为字符串
strconv.Atoi(s string) (i int, err error) 将字符串转换为 int 型
strconv.ParseFloat(s string, bitSize int) (f float64, err error) 将字符串转换为 float64 型

 */
func main()  {
	var orig string = "666"
	var an int
	var newS string

	fmt.Printf("The size of ints is: %d \n", strconv.IntSize)

	an, _ = strconv.Atoi(orig)

	fmt.Printf("The integer is : %d\n", an)

	an = an + 5

	newS = strconv.Itoa(an)

	fmt.Printf("The new string is: %s\n", newS)

}

/*
The size of ints is: 64
The integer is : 666
The new string is: 671
 */