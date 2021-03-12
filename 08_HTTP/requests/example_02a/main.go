package main

import (
    "bufio"
    "fmt"
    "net/http"
    "net/url"
)

func main() {

    target := "https://httpbin.org/post"

    resp, err := http.PostForm(target,
        url.Values{"user": {"john"}, "email": {"john@gmail.com"}})
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
