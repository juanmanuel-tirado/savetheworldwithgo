package main

import "fmt"

func generator(ch chan int) {
    for i:=0;i<5;i++{
        ch <- i
    }
    close(ch)
}

func main() {
    ch := make(chan int)

    go generator(ch)

    for x := range(ch) {
        fmt.Println(x)
    }
    fmt.Println("Done")
}
