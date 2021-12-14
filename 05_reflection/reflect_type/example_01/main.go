package main

import (
    "fmt"
    "reflect"
)

func main() {
    var a int32 = 42
    var b string = "forty two"

    typeA := reflect.TypeOf(a)
    fmt.Println(typeA)

    typeB := reflect.TypeOf(b)
    fmt.Println(typeB)
}
