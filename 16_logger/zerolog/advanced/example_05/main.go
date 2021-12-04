package main

import (
    "net/http"
    "github.com/rs/zerolog"
    "github.com/rs/zerolog/hlog"
    "os"
)

var log zerolog.Logger = zerolog.New(os.Stdout).With().
    Str("app","example_05").Logger()

type MyHandler struct {}

func(c MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    hlog.FromRequest(r).Info().Msg("")
    w.Write([]byte("Perfect!!!"))
    return
}

type Wrapper struct {
    layers []func(http.Handler) http.Handler
}

func NewWrapper() *Wrapper {
    layers := []func(http.Handler) http.Handler {
        hlog.NewHandler(log),
        hlog.RemoteAddrHandler("ip"),
        hlog.UserAgentHandler("user_agent"),
        hlog.RequestIDHandler("req_id", "Request-Id"),
        hlog.MethodHandler("method"),
        hlog.RequestHandler("url"),
    }
    return &Wrapper{layers}
}

func(w *Wrapper) GetWrapper(h http.Handler,i int) http.Handler {
    if i >= len(w.layers) {
        return h
    }
    return w.layers[i](w.GetWrapper(h,i+1))
}

func main() {
    mine := MyHandler{}
    wrapper := NewWrapper()
    h := wrapper.GetWrapper(mine,0)

    panic(http.ListenAndServe(":8090", h))
}
