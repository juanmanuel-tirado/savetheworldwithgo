package main

import (
    "bufio"
    "bytes"
    "fmt"
    "net/http"
    "time"
)

func main() {
    bodyRequest := []byte(`{"user": "john","email":"john@gmail.com"}`)
    bufferBody := bytes.NewBuffer(bodyRequest)

    url := "https://httpbin.org/put"

    header := http.Header{}
    header.Add("Content-type", "application/json")
    header.Add("X-Custom-Header", "somevalue")
    header.Add("User-Agent", "safe-the-world-with-go")

    req, err := http.NewRequest(http.MethodPut, url, bufferBody)

    if err != nil {
        panic(err)
    }

    req.Header = header

    client := http.Client{
        Timeout: time.Second * 5,
        }

    resp, err := client.Do(req)
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
