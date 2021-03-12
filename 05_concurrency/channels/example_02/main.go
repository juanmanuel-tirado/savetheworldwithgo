package main

import (
    "fmt"
    "time"
)

func MrA(ch chan string) {
    time.Sleep(time.Millisecond*500)
    ch <- "This is MrA"
}

func MrB(ch chan string) {
    time.Sleep(time.Millisecond*200)
    ch <- "This is MrB"
}

func main() {
    //ch := make(chan string)
    ch := make(chan string,1)

    ch <- "This is main"

    go MrA(ch)
    go MrB(ch)

    fmt.Println(<-ch)
    fmt.Println(<-ch)
    fmt.Println(<-ch)
}
