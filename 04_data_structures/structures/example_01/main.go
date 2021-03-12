package main

import "fmt"

type Rectangle struct{
    Height int
    Width  int
}

func main() {
    a := Rectangle{}
    fmt.Println(a)

    b := Rectangle{4,4}
    fmt.Println(b)

    c := Rectangle{Width: 10, Height: 3}
    fmt.Println(c)

    d := Rectangle{Width: 7}
    fmt.Println(d)
}
