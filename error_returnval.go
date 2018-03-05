package main

import (
	"math"
	"fmt"
	"errors"
)

func main()  {
	ret, err := MySqrt(16)
	if err != nil {
		fmt.Println("Error! Return values are: ", ret, err)
	} else {
		fmt.Println("It's ok! Return values are: ", ret, err)
	}
	if ret2, err2 := MySqrt(5); err2 != nil {
		fmt.Println("Error! Return values are: ", ret2, err2)
	} else {
		fmt.Println("It's ok! Return values are: ", ret2, err2)
	}
	// named return variables:
	fmt.Println(MySqrt2(5))
}

func MySqrt(a float64) (float64, error) {
	if a < 0 {
		return float64(math.NaN()), errors.New("I won't be able to do a sqrt of negative number!")
	}
	return math.Sqrt(a), nil
}

func MySqrt2(a float64) (ret float64, err error) {
	if a < 0 {
		//then you can use those variables in code
		ret = float64(math.NaN())
		err = errors.New("I won't be able to do a sqrt of negative number!")
	} else {
		ret = math.Sqrt(a)
		//err is not assigned, so it gets default value nil
	}
	//automatically return the named return variables ret and err
	return
}
/**
It's ok! Return values are:  4 <nil>
It's ok! Return values are:  2.23606797749979 <nil>
2.23606797749979 <nil>
 */