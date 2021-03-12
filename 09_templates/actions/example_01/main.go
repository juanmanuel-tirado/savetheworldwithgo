package main

import (
    "text/template"
    "os"
)

type User struct{
    Name string
    Female bool
}

const Msg = `
{{if .Female}}Mrs.{{- else}}Mr.{{- end}} {{.Name}},
Your package is ready.
Thanks,
`

func main() {
    u1 := User{"John", false}
    u2 := User{"Mary", true}

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

