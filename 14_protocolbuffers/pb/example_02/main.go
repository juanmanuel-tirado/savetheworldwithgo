package main

import (
    "github.com/juanmanuel-tirado/savetheworldwithgo/12_protocolbuffers/pb/example_02/user"
    "google.golang.org/protobuf/proto"
    "fmt"
)

func main() {
    userA := user.User{UserId: "John", Email: "john@gmail.com"}
    userB := user.User{UserId: "Mary", Email:"mary@gmail.com"}

    g := user.Group{Id: 1,
        Score: 42.0,
        Category: user.Category_DEVELOPER,
        Users: []*user.User{&userA,&userB},
    }
    fmt.Println("To encode:", g.String())

    encoded, err := proto.Marshal(&g)
    if err != nil {
        panic(err)
    }
    recovered := user.Group{}
    err = proto.Unmarshal(encoded, &recovered)
    fmt.Println("Recovered:", recovered.String())
}


