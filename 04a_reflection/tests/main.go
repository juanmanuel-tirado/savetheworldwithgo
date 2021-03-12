package main

import (
    "reflect"
)

func main() {
    var a int32 = 42
    v := reflect.ValueOf(a)
    v.SetInt(16) // <-- panic
}
