package main

import (
    "bytes"
    "encoding/json"
    "fmt"
)

type User struct {
    UserId string `json:"userId,omitempty"`
    Score int `json:"score,omitempty"`
    password string `json:"password,omitempty"`
}

func main() {

    userA := User{"Gopher", 1000, "admin"}
    userB := User{"BigJ", 10, "1234"}
    userC := User{UserId: "GGBoom", password: "1111"}

    db := []User{userA, userB, userC}
    dbJson, err := json.Marshal(&db)
    if err != nil {
        panic(err)
    }

    fmt.Println(string(dbJson))

    var indented bytes.Buffer
    err = json.Indent(&indented, dbJson,"","    ")
    if err != nil {
        panic(err)
    }
    fmt.Println(indented.String())

    var recovered []User
    err = json.Unmarshal(dbJson, &recovered)
    if err != nil{
        panic(err)
    }

    fmt.Println(recovered)
}
