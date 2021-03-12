package main

import (
    "fmt"
    "google.golang.org/protobuf/proto"
    "github.com/juanmanuel-tirado/savetheworldwithgo/12_protocolbuffers/pb/example_07/user"
)

func main() {
    userA := user.User{UserId: "John", Email: "john@gmail.com"}
    userB := user.User{UserId: "Mary", Email: "mary@gmail.com"}

    teams := map[string]*user.UserList {
        "teamA": &user.UserList{Users:[]*user.User{&userA,&userB}},
        "teamB": nil,
    }

    teamsPB := user.Teams{Teams: teams}

    fmt.Println("To encode:", teamsPB.String())

    encoded, err := proto.Marshal(&teamsPB)
    if err != nil {
        panic(err)
    }
    recovered := user.Teams{}
    err = proto.Unmarshal(encoded, &recovered)
    if err != nil {
        panic(err)
    }
    fmt.Println("Recovered:", recovered.String())
}


