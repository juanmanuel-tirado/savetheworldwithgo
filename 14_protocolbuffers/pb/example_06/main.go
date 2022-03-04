package main

import (
	"fmt"

	"github.com/juanmanuel-tirado/savetheworldwithgo/14_protocolbuffers/pb/example_06/user"
	"google.golang.org/protobuf/proto"
)

func main() {
	goDeveloper := user.Developer{Language: "go"}
	userA := user.User{UserId: "John", Email: "john@gmail.com",
		Type: &user.User_Developer{&goDeveloper}}
	aksOperator := user.Operator{Platform: "aks"}
	userB := user.User{UserId: "Mary", Email: "mary@gmail.com",
		Type: &user.User_Operator{&aksOperator}}

	encodedA, err := proto.Marshal(&userA)
	if err != nil {
		panic(err)
	}
	encodedB, err := proto.Marshal(&userB)
	if err != nil {
		panic(err)
	}

	recoveredA, recoveredB := user.User{}, user.User{}
	err = proto.Unmarshal(encodedA, &recoveredA)
	if err != nil {
		panic(err)
	}
	err = proto.Unmarshal(encodedB, &recoveredB)
	if err != nil {
		panic(err)
	}
	fmt.Println("RecoveredA:", recoveredA.String())
	fmt.Println("RecoveredB:", recoveredB.String())
}
