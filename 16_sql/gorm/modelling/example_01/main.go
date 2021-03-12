package main

import (
    "database/sql/driver"
    "encoding/json"
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

type User struct {
    Name string
    Email string
}

func(u *User) Scan(src interface{}) error {
    input := src.([]byte)
    json.Unmarshal(input,u)
    return nil
}
func(u User) Value()(driver.Value, error) {
    enc, err := json.Marshal(u)
    return enc,err
}

type Operator struct {
    ID uint
    User User `gorm:"embedded,embeddedPrefix:user_"`
    Platform string `gorm:"not null"`
    Dedication uint `gorm:"check:dedication>5"`
}

func main() {
    db, err := gorm.Open(sqlite.Open("/tmp/example01.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    err = db.AutoMigrate(&Operator{})
    if err != nil {
        panic(err)
    }

    op := Operator{
        User: User{
            Name:  "John",
            Email: "john@gmail.com",
        },
        Platform: "k8s",Dedication:10,
    }
    db.Create(&op)
    fmt.Println("Created", op)
    var recovered Operator
    db.First(&recovered,1)
    fmt.Println("Recovered", recovered)
}

