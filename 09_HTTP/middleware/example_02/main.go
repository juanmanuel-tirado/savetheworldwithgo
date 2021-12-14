package main

import "net/http"

type MyHandler struct{}

func (mh *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Perfect!!!"))
    return
}

type Middleware func(http.Handler) http.Handler

func ApplyMiddleware(h http.Handler, middleware ... Middleware) http.Handler {
    for _, next := range middleware {
        h = next(h)
    }
    return h
}

func SimpleMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
        value := w.Header().Get("simple")
        if value == "" {
            value = "X"
        } else {
            value = value + "X"
        }
        w.Header().Set("simple", value)
        next.ServeHTTP(w,r)
    })
}

func main() {

    h := &MyHandler{}

    http.Handle("/three", ApplyMiddleware(
        h, SimpleMiddleware, SimpleMiddleware, SimpleMiddleware))
    http.Handle("/one", ApplyMiddleware(
        h, SimpleMiddleware))

    panic(http.ListenAndServe(":8090", nil))
}
