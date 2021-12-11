package main

import (
    "fmt"
    "reflect"
)

const (
    Pi = 3.14
    Avogadro float32 = 6.022e23
)


func main() {
    fmt.Println("What is the value of Pi? Pi is", Pi)
    fmt.Println(reflect.TypeOf(Pi))
    fmt.Println("Avogadro's Number value is", Avogadro)
    fmt.Println(reflect.TypeOf(Avogadro))
}
