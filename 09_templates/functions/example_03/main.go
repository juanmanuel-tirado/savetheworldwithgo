package main

import (
    "os"
    "text/template"
)

const Header = `
{{block "hello" .}}Hello and welcome{{end}}`

const Welcome = `
{{define "hello"}}
{{range .}}{{print .}} {{end}}
{{end}}
`

func main() {
    musketeers := []string{"Athos", "Porthos", "Aramis","D'Artagnan"}

    helloMsg, err := template.New("start").Parse(Header)
    if err != nil {
        panic(err)
    }

    welcomeMsg, err := template.Must(helloMsg.Clone()).Parse(Welcome)
    if err != nil {
        panic(err)
    }

    if err := helloMsg.Execute(os.Stdout, musketeers); err != nil {
        panic(err)
    }

    if err := welcomeMsg.Execute(os.Stdout, musketeers); err != nil {
        panic(err)
    }


}
