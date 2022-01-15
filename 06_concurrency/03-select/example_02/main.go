package main

import "fmt"

func main() {
    ch := make(chan int)
    //ch := make(chan int,1)

    select {
    case i := <-ch:
        fmt.Println("Received", i)
    default:
        fmt.Println("Nothing received")
    }

    select {
    case ch <- 42:
        fmt.Println("Send 42")
    default:
        fmt.Println("Nothing sent")
    }

    select {
    case i := <-ch:
        fmt.Println("Received", i)
    default:
        fmt.Println("Nothing received")
    }
}
