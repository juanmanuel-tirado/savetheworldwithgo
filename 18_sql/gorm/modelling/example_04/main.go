package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)
// Has many
type User struct {
    ID uint
    Name string
    Email string
    Laptops []Laptop
}

type Laptop struct {
    ID uint
    SerialNumber string
    UserID uint
}

func main() {
    // SQLite does not support foreign key constraints
    db, err := gorm.Open(sqlite.Open("/tmp/example04.db"),
        &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true,})

    if err != nil {
        panic("failed to connect database")
    }

    err = db.AutoMigrate(&User{},&Laptop{})
    if err != nil {
        panic(err)
    }

    laptops := []Laptop{{SerialNumber: "sn0000001"},{SerialNumber: "sn0000002"}}
    u := User{
        Name:    "John",
        Email:   "john@gmail.com",
        Laptops: laptops,
    }
    db.Create(&u)
    fmt.Println("Created", u)
    var recovered User
    db.Preload("Laptops").Find(&recovered,1)
    fmt.Println("Recovered", recovered)
}
