package main

import (
	"fmt"

	"github.com/juanmanuel-tirado/savetheworldwithgo/14_protocolbuffers/pb/example_04/group"
	"google.golang.org/protobuf/proto"
)

func main() {
	userA := group.Group_User{UserId: "John", Email: "john@gmail.com"}
	userB := group.Group_User{UserId: "Mary", Email: "mary@gmail.com"}

	g := group.Group{Id: 1,
		Score:    42.0,
		Category: group.Category_DEVELOPER,
		Users:    []*group.Group_User{&userA, &userB},
	}
	fmt.Println("To encode:", g.String())

	encoded, err := proto.Marshal(&g)
	if err != nil {
		panic(err)
	}
	recovered := group.Group{}
	err = proto.Unmarshal(encoded, &recovered)
	fmt.Println("Recovered:", recovered.String())
}
