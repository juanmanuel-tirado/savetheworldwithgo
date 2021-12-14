package main

import (
    "context"
    "database/sql"
    "fmt"
    "time"
    _ "github.com/mattn/go-sqlite3"
)

func main() {

    db, err := sql.Open("sqlite3", "/tmp/example.db")

    if err != nil {
        fmt.Println(err)
        panic(err)
    }
    defer db.Close()
    db.SetConnMaxLifetime(time.Minute * 3)
    db.SetMaxOpenConns(10)
    db.SetMaxIdleConns(10)

    ctx := context.Background()
    if err := db.PingContext(ctx); err == nil {
        fmt.Println("Database responds")
    } else {
        fmt.Println("Database does not respond")
        panic(err)
    }
    runStatement(db)
}

type User struct {
    Name string
    Email string
}

func runStatement(db *sql.DB) {
    stmt, err := db.Prepare("select * from users where name = ?")
    if err != nil {
        panic(err)
    }
    defer stmt.Close()
    result := stmt.QueryRow("Peter")
    u := User{}

    err =  result.Scan(&u.Name, &u.Email)
    switch {
    case err == sql.ErrNoRows:
        panic(err)
    case err != nil:
        panic(err)
    default:
        fmt.Printf("%v\n",u)
    }
}
