package main

import "fmt"

func a (i int){
    i = 0
}

func b (i *int) {
    *i = 0
}

func main() {
    x := 100

    a(x)
    fmt.Println(x)
    b(&x)
    fmt.Println(x)

    fmt.Println(&x)
}
