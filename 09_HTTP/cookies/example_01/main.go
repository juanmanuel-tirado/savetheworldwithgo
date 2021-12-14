package main

import (
    "fmt"
    "net/http"
    "strconv"
    "time"
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
    return
}

func main() {
    http.HandleFunc("/cookie", cookieSetter)
    go http.ListenAndServe(":8090", nil)

    url := "http://:8090/cookie"
    req, err := http.NewRequest("GET",url,nil)
    if err != nil {
        panic(err)
    }

    client := http.Client{}

    c := http.Cookie{
        Name:"counter", Value:"1", Domain: "127.0.0.1",
        Path: "/", Expires: time.Now().AddDate(1,0,0)}
    req.AddCookie(&c)

    fmt.Println("-->", req.Header)

    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }

    fmt.Println("<--",resp.Header)
}
