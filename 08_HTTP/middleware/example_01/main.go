package main

import (
    "encoding/base64"
    "net/http"
    "strings"
    "fmt"
)

type MyHandler struct{}

func (mh *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Perfect!!!"))
    return
}

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
        header := r.Header.Get("Authorization")
        if header == "" {
            w.WriteHeader(http.StatusUnauthorized)
            return
        }

        authType := strings.Split(header," ")
        fmt.Println(authType)
        if len(authType) != 2 || authType[0] != "Basic" {
            w.WriteHeader(http.StatusUnauthorized)
            return
        }
        credentials,err := base64.StdEncoding.DecodeString(authType[1])
        if err != nil {
            w.WriteHeader(http.StatusUnauthorized)
            return
        }
        if string(credentials) == "Open Sesame" {
            next.ServeHTTP(w, r)
        }
    })
}

func main() {
    targetHandler := MyHandler{}
    panic(http.ListenAndServe(":8090",AuthMiddleware(&targetHandler)))
}
