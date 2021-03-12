package main

import "fmt"

func main() {

    aux := []interface{}{42, "hello", true}

    for _, i := range aux {
        switch t := i.(type) {
        default:
            fmt.Printf("%T --> %s\n", t, i)
        case int:
            fmt.Printf("%T --> %d\n", t, i)
        case string:
            fmt.Printf("%T --> %s\n", t, i)
        case bool:
            fmt.Printf("%T --> %v\n", t, i)
        }
    }
}
