package main

import (
    "fmt"
    "gopkg.in/yaml.v2"
)

type User struct {
    UserId string `yaml:"userId,omitempty"`
    Score int `yaml:"score,omitempty"`
    password string `yaml:"password,omitempty"`
}

func main() {

    userA := User{"Gopher", 1000, "admin"}
    userB := User{"BigJ", 10, "1234"}
    userC := User{UserId: "GGBoom", password: "1111"}

    db := []User{userA, userB, userC}
    dbYaml, err := yaml.Marshal(&db)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(dbYaml))

    var recovered []User
    err = yaml.Unmarshal(dbYaml, &recovered)
    if err != nil{
        panic(err)
    }
    fmt.Println(recovered)
}
