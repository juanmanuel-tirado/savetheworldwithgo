package example_01

import (
    "fmt"
    "strings"
)

type User struct {
    UserId string
    Friends []User
}

func (u *User) GetUserId() string {
    return strings.ToUpper(u.UserId)
}

func (u *User) CountFriends() int {
    return len(u.Friends)
}

func CommonFriend(a *User, b *User) *User {
    for _, af := range a.Friends {
        for _, bf := range b.Friends {
            if af.UserId == bf.UserId {
                return &af
            }
        }
    }
    return nil
}

func ExampleUser() {
    j := User{"John", nil}
    m := User{"Mary", []User{j}}
    fmt.Println(m)
    // Output:
    // {Mary [{John []}]}
}

func ExampleCommonFriend() {
    a := User{"a", nil}
    b := User{"b", []User{a}}
    c := User{"c", []User{a,b}}

    fmt.Println(CommonFriend(&b,&c))
    // Output:
    // &{a []}
}

func ExampleUser_GetUserId() {
    u := User{"John",nil}
    fmt.Println(u.GetUserId())
    // Output:
    // JOHN
}

func ExampleUser_CountFriends() {
    u := User{"John", nil}
    fmt.Println(u.CountFriends())
    // Output:
    // 0
}
