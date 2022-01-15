package main

import (
    "fmt"
    "reflect"
)

type T struct {
    A int32
    B string
    C float32
}

func main() {
    t := T{42, "forty two", 3.14}

    valueT := reflect.ValueOf(t)
    fmt.Println(valueT.Kind(), valueT)

    for i:=0;i<valueT.NumField();i++{
        field := valueT.Field(i)
        fmt.Println(field.Kind(), field.String(), field.Interface())
    }
}


