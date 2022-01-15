package main

import "fmt"

type Rectangle struct{
    Height int
    Width  int
}

func (r Rectangle) Surface() int {
    return r.Height * r.Width
}

func main() {
    r := Rectangle{2,3}
    fmt.Printf("rectangle %v has surface %d",r, r.Surface())
}
