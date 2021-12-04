package main

import (
    "context"
    "fmt"
    "sync/atomic"
    "time"
)

func accum(c *uint32, ctx context.Context) {
    t := time.NewTicker(time.Millisecond*250)
    for {
        select {
        case <- t.C:
            atomic.AddUint32(c, 1)
        case <- ctx.Done():
            fmt.Println("Done context")
            return
        }
    }
}

func main() {
    d := time.Now().Add(time.Second*3)
    ctx, cancel := context.WithDeadline(context.Background(), d)
    defer cancel()

    var counter uint32 = 0

    for i:=0;i<5;i++ {
        go accum(&counter, ctx)
    }

    <- ctx.Done()
    fmt.Println("counter is:", counter)
}
