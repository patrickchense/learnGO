package main

const n  = 5
func main()  {
	buf := [10]int{1,2,3,4,5}
	s, b := buf[:5], buf[5:]
	println(s)
	println(b)
}
/*
[5/10]0xc420041f20
[5/5]0xc420041f48
 */