package main

import "fmt"

type Rectangle struct {
    Height int
    Width int
}

func (r Rectangle) Surface() int {
    return r.Height * r.Width
}

type Box struct {
    Rectangle
    depth int
}

func (b Box) Volume() int {
    return b.Surface() * b.depth
}

func main() {
    b := Box{Rectangle{3,3}, 3}
    fmt.Printf("%+v\n",b)
    fmt.Println("Volume", b.Volume())
}
