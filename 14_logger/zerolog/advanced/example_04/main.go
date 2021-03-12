package main

import (
    "net/http"
    "github.com/rs/zerolog"
    "github.com/rs/zerolog/hlog"
    "os"
)

var log zerolog.Logger = zerolog.New(os.Stdout).With().
    Str("app","example_04").Logger()

type MyHandler struct {}

func(c MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    log.Info().Msg("This is not a request contextual logger")
    hlog.FromRequest(r).Info().Msg("")
    w.Write([]byte("Perfect!!!"))
    return
}

func main() {
    mine := MyHandler{}
    a := hlog.NewHandler(log)
    b := hlog.RemoteAddrHandler("ip")
    c := hlog.UserAgentHandler("user_agent")
    d := hlog.RequestIDHandler("req_id", "Request-Id")

    panic(http.ListenAndServe(":8090", a(b(c(d(mine))))))
}
