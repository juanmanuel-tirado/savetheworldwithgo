package main

import (
    "fmt"
    "time"
)

func sender(out chan int) {
    for i:=0;i<5;i++ {
        time.Sleep(time.Millisecond * 500)
        out <- i
    }
    close(out)
    fmt.Println("sender finished")
}

func main() {

    ch := make(chan int)

    go sender(ch)

    for {
        num, found := <- ch
        if found {
            fmt.Println(num)
        } else {
            fmt.Println("finished")
            break
        }
    }
}
