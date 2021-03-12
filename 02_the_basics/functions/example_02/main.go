package main

import "fmt"

func ops(a int, b int) (int, int) {
    return a + b, a - b
}

func main() {
    sum, subs := ops(2,2)
    fmt.Println("2+2=",sum, "2-2=",subs)
    b, _ := ops(10,2)
    fmt.Println("10+2=",b)
}
