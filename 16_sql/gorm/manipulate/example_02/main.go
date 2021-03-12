package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

type User struct {
    ID uint
    Name string
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
    users := []User{{Name: "John"},{Name: "Mary"},{Name: "Peter"},{Name:"Jeremy"}}
    db.CreateInBatches(users,4)

    var u User
    db.First(&u)
    fmt.Println("First",u)
    u=User{}
    db.Take(&u)
    fmt.Println("Take",u)
    u=User{}
    db.Last(&u)
    fmt.Println("Last",u)
    u=User{}
    db.First(&u,2)
    fmt.Println("First ID=2",u)
    var retrievedUsers []User
    db.Find(&retrievedUsers)
    fmt.Println("Find",retrievedUsers)
    db.Find(&retrievedUsers,[]int{2,4})
    fmt.Println("Find ID=2,ID=4",retrievedUsers)
    u=User{}
    db.Where("name = ?","Jeremy").First(&u)
    fmt.Println("Where name=Jeremy",u)
    db.Where("name LIKE ?","%J%").Find(&retrievedUsers)
    fmt.Println("Where name=%J%",retrievedUsers)
    db.Where("name LIKE ?","%J%").Or("name LIKE ?","%y").Find(&retrievedUsers)
    fmt.Println("Name with J or y",retrievedUsers)
    u=User{}
    db.Where(&User{Name: "Mary"}).First(&u)
    fmt.Println("User with name Mary",u)
    db.Order("name asc").Find(&retrievedUsers)
    fmt.Println("All users ordered by name",retrievedUsers)
}
