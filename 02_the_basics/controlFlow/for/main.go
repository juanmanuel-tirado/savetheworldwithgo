package main

import (
    "fmt"

)

func main() {

    x := 5

    counter := x

    for counter > 0 {
        fmt.Println(counter)
        counter--
    }

    for i:=0; i < x; i++ {
        fmt.Print(i)
    }
    fmt.Println()

    for {
        if x % 2 != 0 {
            fmt.Printf("%d is odd\n", x)
            x++
            continue
        }
        break
    }

    for {
        fmt.Println("Never stop")
        break
    }

}
