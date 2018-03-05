package main

import (
	"fmt"
	"log"
	"io"
)

func main()  {
	defera()
	deferf()
	fmt.Println()
	fmt.Println("tracing example:")
	defer_b()
	fmt.Println()
	fmt.Println("lamba defer example:")
	func1("Go")
}
//使用 defer 的语句同样可以接受参数，下面这个例子就会在执行 defer 语句时打印 0
func defera() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}
//当有多个 defer 行为被注册时，它们会以逆序执行（类似栈，即后进先出）：
func deferf() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}

/*
0
4 3 2 1 0
 */

 //可以用于tracing 函数
func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func defer_a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func defer_b() {
	defer un(trace("b"))
	fmt.Println("in b")
	defer_a()
}

/*
tracing example:
entering: b
in b
entering: a
in a
leaving: a
leaving: b
 */

 //通过lamba函数
func func1(s string) (n int, err error) {
	defer func() {
		log.Printf("func1(%q) = %d, %v", s, n, err)
	}()
	return 7, io.EOF
}
/*
2018/03/05 16:24:55 func1("Go") = 7, EOF
 */