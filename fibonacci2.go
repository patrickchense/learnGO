package main

/*
闭包的属性值被记录下来，外面的a，b会增加
 */
func fib() func() int {
	a, b := 1, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}

func main() {
	f := fib()
	// Function calls are evaluated left-to-right.
	// println(f(), f(), f(), f(), f())
	for i := 0; i <= 9; i++ {
		println(i+2, f())
	}
}

/*
2 2
3 3
4 5
5 8
6 13
7 21
8 34
9 55
10 89
11 144
 */