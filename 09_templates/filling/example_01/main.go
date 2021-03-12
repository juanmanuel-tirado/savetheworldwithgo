package main

import (
    "text/template"
    "os"
)

type User struct{
    Name string
    UserId string
    Email string
}

const Msg = `Dear {{.Name}},
You were registered with id {{.UserId}}
and e-mail {{.Email}}.
`

func main() {
    u := User{"John", "John33", "john@gmail.com"}
    t := template.Must(template.New("msg").Parse(Msg))
    err := t.Execute(os.Stdout, u)
    if err != nil {
        panic(err)
    }
}
