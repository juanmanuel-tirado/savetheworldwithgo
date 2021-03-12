package main

import (
    "fmt"
    "reflect"
)

type T struct {
    A int32
    B string
}

func main() {
    t := T{42, "forty two"}

    typeT := reflect.TypeOf(t)
    fmt.Println(typeT)

    for i:=0;i<typeT.NumField();i++{
        field := typeT.Field(i)
        fmt.Println(field.Name,field.Type)
    }
}
