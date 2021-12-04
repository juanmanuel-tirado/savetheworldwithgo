package main

import (
    "fmt"
    "sync/atomic"
    "time"
)

func increaser(counter *int32) {
    for {
        atomic.AddInt32(counter,2)
        time.Sleep(time.Millisecond*500)
    }
}

func decreaser(counter *int32) {
    for {
        atomic.AddInt32(counter, -1)
        time.Sleep(time.Millisecond*250)
    }
}

func main() {
    var counter int32 = 0

    go increaser(&counter)
    go decreaser(&counter)

    for i:=0;i<5;i++{
        time.Sleep(time.Millisecond*500)
        fmt.Println(atomic.LoadInt32(&counter))
    }
    fmt.Println(atomic.LoadInt32(&counter))
}
