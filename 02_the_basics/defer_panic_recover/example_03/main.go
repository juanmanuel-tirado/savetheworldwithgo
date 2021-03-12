package main

import "fmt"

func something() {
    defer func() {
        r := recover()
        if r != nil{
            fmt.Println("No need to panic if i=",r)
        }
    }()
    for i:=0;i<5;i++ {
        fmt.Println(i)
        if i > 2 {
            panic(i)
        }
    }
    fmt.Println("Closed something  normally")
}

func main () {
    defer fmt.Println("closed main")

    something()
    fmt.Println("Main was finished")
}
