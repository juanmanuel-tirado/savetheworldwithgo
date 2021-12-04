package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)
// Many to many
type User struct {
    ID uint
    Name string
    Email string
    Languages []Language `gorm:"many2many:user_languages"`
    Laptops []*Laptop `gorm:"many2many:user_laptops"`
}

type Laptop struct {
    ID uint
    SerialNumber string
    Users []*User `gorm:"many2many:user_laptops"`
}

type Language struct {
    ID uint
    Name string
}

func main() {
    db, err := gorm.Open(sqlite.Open("/tmp/example05.db"), &gorm.Config{})

    if err != nil {
        panic("failed to connect database")
    }

    err = db.AutoMigrate(&Language{},&User{},&Laptop{})
    if err != nil {
        panic(err)
    }

    laptops := []*Laptop{{SerialNumber: "sn0000001"},{SerialNumber: "sn0000002"}}
    languages := []Language{{Name: "Go"},{Name: "Python"}}
    u := User{
        Name:    "John",
        Email:   "john@gmail.com",
        Laptops: laptops,
        Languages: languages,
    }
    db.Create(&u)
    fmt.Println("Created", u)
    var recovered User
    db.Preload("Laptops").Find(&recovered,1)
    fmt.Println("Recovered", recovered)
}
