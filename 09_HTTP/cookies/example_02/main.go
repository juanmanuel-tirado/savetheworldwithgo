package main

import (
    "fmt"
    "net/http"
    "net/http/cookiejar"
    url2 "net/url"
    "strconv"
)

func cookieSetter(w http.ResponseWriter, r *http.Request) {
    counter, err := r.Cookie("counter")
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    value, err := strconv.Atoi(counter.Value)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    value = value + 1
    newCookie := http.Cookie{
        Name: "counter",
        Value: strconv.Itoa(value),
    }
    http.SetCookie(w, &newCookie)
    w.WriteHeader(http.StatusOK)
}

func main() {
    http.HandleFunc("/cookie", cookieSetter)
    go http.ListenAndServe(":8090", nil)

    jar, err := cookiejar.New(nil)
    if err != nil {
        panic(err)
    }
    cookies := []*http.Cookie{
        &http.Cookie{Name:"counter",Value:"1"},
    }

    url := "http://localhost:8090/cookie"
    u, _ := url2.Parse(url)
    jar.SetCookies(u, cookies)

    client := http.Client{Jar: jar}

    for i:=0; i<5; i++ {
        _, err := client.Get(url)
        if err != nil {
            panic(err)
        }
        fmt.Println("Client cookie",jar.Cookies(u))
    }
}
