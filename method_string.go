package main

import (
	"fmt"
	"strconv"
)

type TwoInts1 struct {
	a int
	b int
}

func main() {
	two1 := new(TwoInts1)
	two1.a = 12
	two1.b = 10
	fmt.Printf("two1 is: %v\n", two1)
	fmt.Println("two1 is:", two1)
	fmt.Printf("two1 is: %T\n", two1)
	fmt.Printf("two1 is: %#v\n", two1)
}
/*
当你广泛使用一个自定义类型时，最好为它定义 String()方法。从上面的例子也可以看到，格式化描述符 %T 会给出类型的完全规格，%#v 会给出实例的完整输出，包括它的字段（在程序自动生成 Go 代码时也很有用）
 */
func (tn *TwoInts1) String() string {
	return "(" + strconv.Itoa(tn.a) + "/" + strconv.Itoa(tn.b) + ")"
}
/*
two1 is: (12/10)
two1 is: (12/10)
two1 is: *main.TwoInts1
two1 is: &main.TwoInts1{a:12, b:10}
 */