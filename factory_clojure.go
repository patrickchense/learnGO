package main

import "strings"

func main()  {
	//new func
	addBmp := MakeAddSuffix(".bmp")
	addJpeg := MakeAddSuffix(".jpeg")

	//use
	println(addBmp("file"))
	println(addJpeg("file"))
}

func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
/**
file.bmp
file.jpeg
 */