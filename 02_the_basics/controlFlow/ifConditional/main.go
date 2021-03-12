package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    x := rand.Float32()

    if x < 0.5 {
        fmt.Println("head")
    } else {
        fmt.Println("tail")
    }
}
