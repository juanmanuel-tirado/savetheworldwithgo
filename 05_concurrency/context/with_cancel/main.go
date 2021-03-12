package main

import (
    "context"
    "fmt"
    "sync/atomic"
    "time"
)

func setter(id int, c *int32, ctx context.Context) {
    t := time.NewTicker(time.Millisecond*300)
    for {
        select {
        case <- ctx.Done():
            fmt.Println("Done", id)
            return
        case <- t.C:
            atomic.AddInt32(c, 1)
        }
    }
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())

    var c int32 = 0
    for i:=0;i<5;i++ {
        go setter(i, &c, ctx)
    }

    time.Sleep(time.Second * 1)
    fmt.Println("Final check: ", c)

    cancel()
    time.Sleep(time.Second)
}
