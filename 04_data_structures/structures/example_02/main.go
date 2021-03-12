package main

import "fmt"

type Rectangle struct{
    Height int
    Width  int
}

func NewRectangle(height int, width int) *Rectangle {
    return &Rectangle{height, width}
}

func main() {
    a := Rectangle{Height: 7}
    fmt.Println(a)

    r := NewRectangle(2,3)
    fmt.Println(r)
    fmt.Println(*r)
}
