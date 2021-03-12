package main

import (
    "context"
    "fmt"
    "time"
)

func calc(ctx context.Context) {
    switch ctx.Value("action") {
    case "quick":
        fmt.Println("quick answer")
    case "slow":
        time.Sleep(time.Millisecond*500)
        fmt.Println("slow answer")
    case <- ctx.Done():
        fmt.Println("Done!!!")
    default:
        panic("unknown action")
    }
}

func main() {
    t := time.Millisecond*300
    ctx, cancel := context.WithTimeout(context.Background(), t)
    qCtx := context.WithValue(ctx, "action", "quick")
    defer cancel()

    go calc(qCtx)
    <-qCtx.Done()

    ctx2, cancel2 := context.WithTimeout(context.Background(), t)
    sCtx := context.WithValue(ctx2, "action", "slow")
    defer cancel2()

    go calc(sCtx)
    <-sCtx.Done()

    fmt.Println("Finished")
}
