package main

import (
    "fmt"
    "time"
)

func reaction(t *time.Ticker) {
    for {
        select {
        case x := <-t.C:
            fmt.Println("quick",x)
        }
    }
}

func slowReaction(t *time.Timer) {
    select {
    case x := <-t.C:
        fmt.Println("slow", x)
    }
}

func main() {
    quick := time.NewTicker(time.Second)
    slow := time.NewTimer(time.Second * 5)
    stopper := time.NewTimer(time.Second * 4)
    go reaction(quick)
    go slowReaction(slow)

    <- stopper.C
    quick.Stop()

    stopped := slow.Stop()
    fmt.Println("Stopped before the event?",stopped)
}
