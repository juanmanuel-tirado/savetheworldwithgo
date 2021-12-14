package main

import "fmt"

func accumulator(increment int) func() int {
    i:=0
    return func() int {
        i = i + increment
        return i
    }
}

func main() {

    a := accumulator(1)
    b := accumulator(2)

    fmt.Println("a","b")
    for i:=0;i<5;i++ {
        fmt.Println(a(),b())
    }
}
