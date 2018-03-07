package main

import "time"

func main()  {
	ch := make(chan string)
	go sendData(ch)
	go getDate(ch)

	time.Sleep(100000)
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
}

func getDate(ch chan string) {
	var input string
	for {
		input = <- ch
		println(input)
	}
}
