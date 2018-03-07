package main

import (
	"crypto/sha1"
	"io"
	"fmt"
	"log"
)
/*
通过调用 sha1.New() 创建了一个新的 hash.Hash 对象，用来计算 SHA1 校验值。Hash 类型实际上是一个接口，它实现了 io.Writer 接口
通过 io.WriteString 或 hasher.Write 将给定的 []byte 附加到当前的 hash.Hash 对象中。
 */
func main() {
	hasher := sha1.New()
	io.WriteString(hasher, "test")
	b := []byte{}
	fmt.Printf("Result: %x\n", hasher.Sum(b))
	fmt.Printf("Result: %d\n", hasher.Sum(b))
	//
	hasher.Reset()
	data := []byte("We shall overcome!")
	n, err := hasher.Write(data)
	if n!=len(data) || err!=nil {
		log.Printf("Hash write error: %v / %v", n, err)
	}
	checksum := hasher.Sum(b)
	fmt.Printf("Result: %x\n", checksum)
}
/*
Result: a94a8fe5ccb19ba61c4c0873d391e987982fbbd3
Result: [169 74 143 229 204 177 155 166 28 76 8 115 211 145 233 135 152 47 187 211]
Result: e2222bfc59850bbb00a722e764a555603bb59b2a
 */