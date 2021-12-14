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

func RunTransaction(u *User, db *gorm.DB) error{
    tx := db.Begin()
    if tx.Error != nil {
        return tx.Error
    }
    if err := tx.Model(u).Update("Email","newemail").Error; err != nil {
        return err
    }
    tx.SavePoint("savepoint")
    if err := tx.Create(u).Error; err != nil{
        tx.RollbackTo("savepoint")
    }
    return tx.Commit().Error
}

func main() {
    db, err := gorm.Open(sqlite.Open("/tmp/example02.db"), &gorm.Config{})

    if err != nil {
        panic("failed to connect database")
    }

    err = db.AutoMigrate(&User{})
    if err != nil {
        panic(err)
    }

    u := User{Name: "John", Email: "john@gmail.com"}
    db.Create(&u)

    err = RunTransaction(&u, db)
    if err != nil {
        fmt.Println(err)
    }
    var retrieved User
    db.First(&retrieved)
    fmt.Println("Retrieved", retrieved)
}
