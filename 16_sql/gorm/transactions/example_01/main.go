package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

type User struct {
    ID uint
    Name string
    Email string
}
func main() {
    db, err := gorm.Open(sqlite.Open("/tmp/example01.db"), &gorm.Config{})

    if err != nil {
        panic("failed to connect database")
    }

    err = db.AutoMigrate(&User{})
    if err != nil {
        panic(err)
    }

    u := User{Name: "John", Email: "john@gmail.com"}
    db.Create(&u)

    db.Transaction(func(tx *gorm.DB) error {
        if err := tx.Model(&u).Update("Email","newemail").Error; err != nil {
            return err
        }
        var inside User
        tx.First(&inside)
        fmt.Println("Retrieved inside transaction", inside)
        if err := tx.Create(&u).Error; err != nil {
            return err
        }
        return nil
    })
    var retrieved User
    db.First(&retrieved)
    fmt.Println("Retrieved", retrieved)
}
