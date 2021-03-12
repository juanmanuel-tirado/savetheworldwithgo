package main

import (
    "encoding/json"
    "fmt"
    "github.com/juanmanuel-tirado/savetheworldwithgo/12_protocolbuffers/pb/example_08/user"
)

func main() {
    userA := user.User{UserId: "John", Email: "john@gmail.com"}
    userB := user.User{UserId: "Mary", Email: "mary@gmail.com"}

    teams := map[string]*user.UserList {
        "teamA": &user.UserList{Users:[]*user.User{&userA,&userB}},
        "teamB": nil,
    }

    teamsPB := user.Teams{Teams: teams}

    encoded, err := json.Marshal(&teamsPB)
    if err != nil {
        panic(err)
    }
    recovered := user.Teams{}
    err = json.Unmarshal(encoded, &recovered)
    if err != nil {
        panic(err)
    }
    fmt.Println("Recovered:", recovered.String())
}


