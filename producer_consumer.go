package main


func main()  {
	ch := make(chan int)
	done := make(chan bool)
	go func(ch chan int) {
		for i:= 0; i< 10 ; i++ {
			ch <- i * 10
		}
		close(ch)  //不用的通道要关闭，不然会导致deadlock
	}(ch)

	go func(ch chan int, done chan bool) {
		for i := range  ch {
			println(i)
		}
		done <- true
	}(ch, done)

	<- done
}

//func numGen(start, count int, out chan<- int) {
//	for i := 0; i < count; i++ {
//		out <- start
//		start = start + count
//	}
//	close(out)
//}
//
//// integer consumer:
//func numEchoRange(in <-chan int, done chan<- bool) {
//	for num := range in {
//		fmt.Printf("%d\n", num)
//	}
//	done <- true
//}
//
//func main() {
//	numChan := make(chan int)
//	done := make(chan bool)
//	go numGen(0, 10, numChan)
//	go numEchoRange(numChan, done)
//
//	<-done
//}