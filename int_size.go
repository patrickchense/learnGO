package main

import (
	"unsafe"
)

func main()  {
	s := unsafe.Sizeof(32)
	println(s)  //8
}
