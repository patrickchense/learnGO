package main

import "fmt"
/*
通过指针，修改值，传递值需要拷贝，指针不用
 */
func Multiply(a, b int, reply *int) {
	*reply = a * b
}

func main() {
	n := 0
	reply := &n
	Multiply(10, 5, reply)
	fmt.Println("Multiply:", *reply) // Multiply: 50
}
//Multiply: 50