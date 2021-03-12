package main

import (
    "context"
    "database/sql"
    "fmt"
    "time"
    _ "github.com/mattn/go-sqlite3"
)

func main() {

    db, err := sql.Open("sqlite3","/tmp/example.db")

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

    get_rows(db)
}

type User struct {
    Name string
    Email string
}

func get_rows(db *sql.DB) {
    ctx := context.Background()
    rows, err := db.QueryContext(ctx, "SELECT * from users")
    if err != nil {
        panic(err)
    }
    defer rows.Close()

    for rows.Next() {
        u := User{}
        rows.Scan(&u.Name,&u.Email)
        fmt.Printf("%v\n",u)
    }
}
