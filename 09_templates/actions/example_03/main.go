package main

import (
    "text/template"
    "os"
)

const msg = `
The musketeers are:
{{range .}}{{print .}} {{end}}
`

func main() {
    musketeers := []string{"Athos", "Porthos", "Aramis","D'Artagnan"}

    t := template.Must(template.New("msg").Parse(msg))
    err := t.Execute(os.Stdout, musketeers)
    if err != nil {
        panic(err)
    }
}

