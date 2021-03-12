package main

import "fmt"

func CloseMsg() {
    fmt.Println("Closed!!!")
}

func main() {
    defer CloseMsg()

    fmt.Println("Doing something...")
    defer fmt.Println("Certainly closed!!!")
    fmt.Println("Doing something else...")

}
