package main

import "fmt"

type Rectangle struct{
    Height int
    Width  int
}

func (r Rectangle) Enlarge(factor int) {
    r.Height = r.Height * factor
    r.Width = r.Width * factor
}

func (r *Rectangle) EnlargeP(factor int) {
    r.Height = r.Height * factor
    r.Width = r.Width * factor
}

func main() {
    rect := Rectangle{2,2}
    fmt.Println(rect)

    rect.Enlarge(2)
    fmt.Println(rect)

    rect.EnlargeP(2)
    fmt.Println(rect)
}
