package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    target := make([]byte,50)
    n, err := os.Stdin.Read(target)
    if err != nil {
        panic(err)
    }
    msg := string(target[:n])
    fmt.Println(n,strings.ToUpper(msg))
}
