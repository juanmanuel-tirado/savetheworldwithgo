package main

import (
    "text/template"
    "os"
)

type User struct{
    Name string
    Score uint32
}

const Msg = `
{{.Name}} your score is {{.Score}}
your level is:
{{if le .Score 50}}Amateur
{{else if le .Score 80}}Professional 
{{else}}Expert
{{end}}
`

func main() {
    u1 := User{"John", 30}
    u2 := User{"Mary", 80}

    t := template.Must(template.New("msg").Parse(Msg))
    err := t.Execute(os.Stdout, u1)
    if err != nil {
        panic(err)
    }
    err = t.Execute(os.Stdout, u2)
    if err != nil {
        panic(err)
    }
}
