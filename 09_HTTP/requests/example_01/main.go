package main

import (
    "bufio"
    "fmt"
    "net/http"
)

func main() {

    resp, err := http.Get("https://httpbin.org/get")
    if err != nil {
        panic(err)
    }

    fmt.Println(resp.Status)
    fmt.Println(resp.Header["Content-Type"])

    defer resp.Body.Close()
    buf := bufio.NewScanner(resp.Body)

    for buf.Scan() {
        fmt.Println(buf.Text())
    }
}
