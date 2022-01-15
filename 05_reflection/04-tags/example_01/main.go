package main

import (
    "fmt"
    "reflect"
)

type User struct {
    UserId string `tagA:"valueA1" tagB: "valueA2"`
    Email string `tagB:"value"`
    Password string `tagC:"v1 v2"`
}

func main() {
    T := reflect.TypeOf(User{})

    fieldUserId := T.Field(0)
    t := fieldUserId.Tag
    fmt.Println("StructTag is:",t)
    v, _ := t.Lookup("tagA")
    fmt.Printf("tagA: %s\n", v)

    fieldEmail, _ := T.FieldByName("Email")
    vEmail := fieldEmail.Tag.Get("tagB")
    fmt.Println("email tagB:", vEmail)

    fieldPassword, _ := T.FieldByName("Password")
    fmt.Printf("Password tags: [%s]\n",fieldPassword.Tag)
    fmt.Println(fieldPassword.Tag.Get("tagC"))
}
