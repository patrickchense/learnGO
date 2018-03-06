package main

import (
	"fmt"
	"reflect"
)

func main()  {
	fv := func () {fmt.Println("Hello world")}
	fv()
	fmt.Printf("func fv type: %s", reflect.TypeOf(fv))
}
/**
Hello world
func fv type: func()
 */