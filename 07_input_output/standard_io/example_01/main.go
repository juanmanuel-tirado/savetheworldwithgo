package main

import (
    "fmt"
    "os"
)

func main() {
    msg := []byte("Save the world with Go!!!\n")
    n, err := os.Stdout.Write(msg)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Written %d characters\n",n)
}
