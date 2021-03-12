package main

import (
    "fmt"
    "reflect"
)

type Circle struct {
    x int
    y int
    radius int
}

func main() {
    ac := struct{x int; y int; radius int}{1,2,3}
    c := Circle{10,10,3}

    fmt.Printf("%+v\n", ac)
    fmt.Println(reflect.TypeOf(ac))
    fmt.Printf("%+v\n",c)
    fmt.Println(reflect.TypeOf(c))

    ac.x=3
    fmt.Printf("%+v\n", ac)

    ac = c
    fmt.Printf("%+v\n", ac)
    fmt.Println(reflect.TypeOf(ac))
}
