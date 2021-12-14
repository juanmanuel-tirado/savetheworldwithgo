package main

import (
    "text/template"
    "os"
)

const Msg = `
The fourth musketeer is:
{{slice . 3}}
`

func main() {
    musketeers := []string{"Athos", "Porthos", "Aramis","D'Artagnan"}

    t := template.Must(template.New("msg").Parse(Msg))

    err := t.Execute(os.Stdout, musketeers)
    if err != nil {
        panic(err)
    }
}
