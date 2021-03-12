package main

import (
    "strings"
    "text/template"
    "os"
)

const Msg = `
The musketeers are:
{{join . ", "}}
`

func main() {
    musketeers := []string{"Athos", "Porthos", "Aramis","D'Artagnan"}

    funcs := template.FuncMap{"join": strings.Join}

    t, err := template.New("msg").Funcs(funcs).Parse(Msg)
    if err != nil {
        panic(err)
    }
    err = t.Execute(os.Stdout, musketeers)
    if err != nil {
        panic(err)
    }
}
