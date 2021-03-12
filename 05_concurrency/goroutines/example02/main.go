package main

import (
    "time"
    "fmt"
)

func ShowIt(t time.Duration, num int){
    for {
        time.Sleep(t)
        fmt.Println(num)
    }
}

func main() {
    go ShowIt(time.Second, 100)
    go ShowIt(time.Millisecond*500,10)
    go ShowIt(time.Millisecond*250,1)

    time.Sleep(time.Millisecond*1200)
}
