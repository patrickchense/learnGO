package main

func main()  {
PrintIntRec(1)
}

func PrintIntRec(n int) {
	if n != 10 {
		PrintIntRec(n + 1)
	}
	println(n)
}
