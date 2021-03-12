package main

import (
    "fmt"
    "reflect"
)

func main() {
    var a int32 = 42
    var b string = "forty two"

    valueOfA := reflect.ValueOf(a)
    fmt.Println(valueOfA.Interface())

    valueOfB := reflect.ValueOf(b)
    fmt.Println(valueOfB.Interface())
}
