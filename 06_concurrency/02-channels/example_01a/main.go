package main

import (
    "fmt"
    "time"
)

func generator(ch chan int) {
    fmt.Println("generator waits for n")
    n := <- ch
    fmt.Println("n is",n)
    sum := 0
    for i:=0;i<n;i++ {
        sum = sum + i
    }
    ch <- sum
}

func main() {

    ch := make(chan int)

    go generator(ch)

    fmt.Println("main waits for result...")
    time.Sleep(time.Second)
    ch <- 5
    result := <- ch

    fmt.Println(result)
}
