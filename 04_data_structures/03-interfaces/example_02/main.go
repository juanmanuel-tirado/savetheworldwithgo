package main

import "fmt"

func main() {
    var aux interface{}

    fmt.Println(aux)

    aux = 10
    fmt.Println(aux)

    aux = "hello"
    fmt.Println(aux)
}
