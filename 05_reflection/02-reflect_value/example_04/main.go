package main

import (
    "fmt"
    "reflect"
    "strings"
)

type T struct {
    A string
    B int32
    C string
}

func main() {
    t := T{"hello",42,"bye"}
    fmt.Println(t)

    valueOfT := reflect.ValueOf(&t).Elem()
    for i:=0; i< valueOfT.NumField(); i++ {
        f := valueOfT.Field(i)
        if f.Kind() == reflect.String {
            current := f.String()
            f.SetString(strings.ToUpper(current))
        }
    }

    fmt.Println(t)
}
