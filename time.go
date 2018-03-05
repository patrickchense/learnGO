package main

import (
	"fmt"
	"time"
)

var week time.Duration

func main()  {
	t := time.Now()
	fmt.Println(t)
	fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())

	t = time.Now().UTC()
	fmt.Println(t)
	fmt.Println(time.Now())

	week = 60 * 60 * 24 * 7 * 1e9
	week_from_now := t.Add(week)
	fmt.Println(week_from_now)

	fmt.Println(t.Format(time.RFC822))
	fmt.Println(t.Format(time.ANSIC))
	fmt.Println(t.Format("02 Jan 2006 15:04"))
	s := t.Format("20060102")
	fmt.Println(t, "=>", s)
}

/*
2018-03-05 14:10:02.467137645 +0800 CST m=+0.000164679
05.03.2018
2018-03-05 06:10:02.467239693 +0000 UTC
2018-03-05 14:10:02.46724253 +0800 CST m=+0.000269530
2018-03-12 06:10:02.467239693 +0000 UTC
05 Mar 18 06:10 UTC
Mon Mar  5 06:10:02 2018
05 Mar 2018 06:10
2018-03-05 06:10:02.467239693 +0000 UTC => 20180305
 */
