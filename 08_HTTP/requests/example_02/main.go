package main

import (
    "bufio"
    "bytes"
    "fmt"
    "net/http"
)

func main() {
    bodyRequest := []byte(`{"user": "john","email":"john@gmail.com"}`)
    bufferBody := bytes.NewBuffer(bodyRequest)

    url := "https://httpbin.org/post"

    resp, err := http.Post(url, "application/json", bufferBody)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println(resp.Status)
    bodyAnswer := bufio.NewScanner(resp.Body)
    for bodyAnswer.Scan() {
        fmt.Println(bodyAnswer.Text())
    }
}
