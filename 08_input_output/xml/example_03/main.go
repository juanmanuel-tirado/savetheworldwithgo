package main

import (
    "encoding/xml"
    "fmt"
)

type User struct {
    UserId string `xml:"userId,omitempty"`
    Score int `xml:"score,omitempty"`
    password string `xml:"password,omitempty"`
}

type UsersArray struct {
    Users []User `xml:"users,omitempty"`
}

func main() {

    userA := User{"Gopher", 1000, "admin"}
    userB := User{"BigJ", 10, "1234"}
    userC := User{UserId: "GGBoom", password: "1111"}

    db := UsersArray{[]User{userA, userB, userC}}
    dbXML, err := xml.Marshal(&db)
    if err != nil {
        panic(err)
    }

    var recovered UsersArray
    err = xml.Unmarshal(dbXML, &recovered)
    if err != nil{
        panic(err)
    }
    fmt.Println(recovered)

    var indented []byte
    indented, err = xml.MarshalIndent(recovered,"","    ")
    if err != nil {
        panic(err)
    }
    fmt.Println(string(indented))
}
