package main

func main() {
	ch := make(chan int)
	go func(a int, b int, ch chan int) {
			c := a + b
			ch <- c
	}(1, 3, ch)

	println(<- ch) //4

}

