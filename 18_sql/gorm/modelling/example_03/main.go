package main

import (
    "fmt"
    "gorm.io/driver/sqlite"

    "gorm.io/gorm"
)
// Has one
type User struct {
    ID uint
    Name string
    Email string
    Laptop Laptop
}

type Laptop struct {
    ID uint
    SerialNumber string
    UserID uint
}

func main() {
    // SQLite does not support foreign key constraints
    db, err := gorm.Open(sqlite.Open("/tmp/example03.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    err = db.AutoMigrate(&User{},&Laptop{})
    if err != nil {
        panic(err)
    }

    u := User{Name: "John", Email: "john@gmail.com",Laptop: Laptop{SerialNumber: "sn101939"}}
    db.Create(&u)
    fmt.Println("Created", u)
    var recovered User
    db.Preload("Laptop").First(&recovered,1)
    fmt.Println("Recovered",recovered)

}

