package main

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

type User struct {
    ID uint
    Name string
    Email string
}

func main() {
    db, err := gorm.Open(sqlite.Open("/tmp/test.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    err = db.AutoMigrate(&User{})
    if err != nil {
        panic(err)
    }

    u := User{Name: "John", Email: "john@gmail.com"}
    db.Create(&u)

    var recovered User
    db.First(&recovered,"name=?","John")
    fmt.Println("Recovered", recovered)

    db.Model(&recovered).Update("Email","newemail")
    db.First(&recovered,1)
    fmt.Println("After update", recovered)

    db.Delete(&recovered, 1)
}
