package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

type User struct {
    ID uint
    Name string
    GroupID uint
    Group Group
}

type Group struct {
    ID uint
    Name string
}

func main() {
    db, err := gorm.Open(sqlite.Open("/tmp/example04.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    err = db.AutoMigrate(&Group{},&User{})
    if err != nil {
        panic(err)
    }

    g := Group{Name: "TheCoolOnes"}
    u := User{Name: "John", Group: g}
    db.Create(&u)

    var recovered User
    db.Preload("Group").First(&recovered,1)
    fmt.Println("Recovered", recovered)
}
