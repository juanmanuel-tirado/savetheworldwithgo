package main

import "fmt"

func something() {
    defer fmt.Println("closed something")
    for i:=0;i<5;i++ {
        fmt.Println(i)
        if i > 2 {
            panic("Panic was called")
        }
    }
}

func main () {
    defer fmt.Println("closed main")
    something()
    fmt.Println("Something was finished")
}
