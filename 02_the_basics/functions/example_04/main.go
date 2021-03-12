package main

import "fmt"

func doit(operator func(int,int) int, a int, b int) int {
    return operator(a,b)
}

func sum(a int, b int) int {
    return a + b
}

func multiply(a int, b int) int {
    return a * b
}

func main() {
    c := doit(sum, 2, 3)
    fmt.Println("2+3=", c)
    d := doit(multiply, 2, 3)
    fmt.Println("2*3=", d)
}
