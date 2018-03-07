package main

import (
	"time"
	"fmt"
)
/*
超时模式: overtime pattern
timeout := make(chan bool, 1)
go func() {
        time.Sleep(1e9) // one second
        timeout <- true
}()

select {
    case <-ch:
        // a read from ch has occured
    case <-timeout:
        // the read from ch has timed out
        break
}

取消长时间的goroutine
ch := make(chan error, 1)
go func() { ch <- client.Call("Service.Method", args, &reply) } ()
select {
case resp := <-ch
    // use resp and reply
case <-time.After(timeoutNs):
    // call timed out
    break
}

先到先返回模式: 比如从多个库读取
func Query(conns []conn, query string) Result {
    ch := make(chan Result, 1)
    for _, conn := range conns {
        go func(c Conn) {
            select {
            case ch <- c.DoQuery(query):
            default:
            }
        }(conn)
    }
    return <- ch
}
 */
func main() {
	tick := time.Tick(1e8)
	boom := time.After(5e8)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(5e7)
		}
	}
}
/*
    .
    .
tick.
    .
    .
tick.
    .
    .
tick.
    .
    .
tick.
    .
    .
tick.
BOOM!
 */
