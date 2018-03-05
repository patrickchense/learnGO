package main

import (
	"math/big"
	"fmt"
)

func main()  {
	jiecheng(30)
}

func jiecheng(n int64) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	}
	a := big.NewInt(n)
	x := a.Mul(a, jiecheng(n - 1))
	fmt.Printf("jiecheng %d = %d", n, x)
	fmt.Println()
	return x
}
/*
jiecheng 1 = 1
jiecheng 2 = 2
jiecheng 3 = 6
jiecheng 4 = 24
jiecheng 5 = 120
jiecheng 6 = 720
jiecheng 7 = 5040
jiecheng 8 = 40320
jiecheng 9 = 362880
jiecheng 10 = 3628800
jiecheng 11 = 39916800
jiecheng 12 = 479001600
jiecheng 13 = 6227020800
jiecheng 14 = 87178291200
jiecheng 15 = 1307674368000
jiecheng 16 = 20922789888000
jiecheng 17 = 355687428096000
jiecheng 18 = 6402373705728000
jiecheng 19 = 121645100408832000
jiecheng 20 = 2432902008176640000
jiecheng 21 = 51090942171709440000
jiecheng 22 = 1124000727777607680000
jiecheng 23 = 25852016738884976640000
jiecheng 24 = 620448401733239439360000
jiecheng 25 = 15511210043330985984000000
jiecheng 26 = 403291461126605635584000000
jiecheng 27 = 10888869450418352160768000000
jiecheng 28 = 304888344611713860501504000000
jiecheng 29 = 8841761993739701954543616000000
jiecheng 30 = 265252859812191058636308480000000
 */