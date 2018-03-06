package main

import "fmt"

func main()  {
	s := ChangeAscii("ab日本語")
	fmt.Printf("change ascii string: %s -> %s", "ab日本語", s)
	fmt.Println()
	s = ChangeAscii("ab日本語bb")
	fmt.Printf("change ascii string: %s -> %s", "ab日本語bb", s)
}

func ChangeAscii(s string) string {
	//replace not ascii character with ?
	// 0 -127 int
	var b []byte
	var nbytes int
	init := false
	for i, runeValue := range s {
		r := isAscii(runeValue)
		if r > 0 {
			if !init {
				continue
			} else {
				nbytes++
				b = append(b, byte(runeValue))
				continue
			}
		}
		if !init {
			b = make([]byte, i)
			copy(b, s[:i])
			nbytes = i
			init = true
		}
		nbytes ++
		b = append(b, '?')
	}
	if b == nil {
		return s
	}
	return string(b[:nbytes])
}

func isAscii(i int32) int32 {
	if i < 128 && i >= 0 {
		return i
	}
	return -1
}

/*
change ascii string: ab日本語 -> ab???
change ascii string: ab日本語bb -> ab???bb
 */