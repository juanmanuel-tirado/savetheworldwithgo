package main

import "fmt"

func main() {

    a := []int{0,1,2,3,4}

    fmt.Println(a, len(a), cap(a))

    b := append(a,5)
    fmt.Println(b, len(b), cap(b))
    b = append(b,6)
    fmt.Println(b, len(b), cap(b))

    c := b[1:4]
    fmt.Println(c, len(c), cap(c))

    d := make([]int,5,10)
    fmt.Println(d, len(d), cap(d))
    // d[6]=5 --> This will fail

}
