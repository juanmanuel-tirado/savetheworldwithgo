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
    create_table(db)
    insert_rows(db)
}

const (
    USERS_TABLE=`CREATE TABLE users(
    name varchar(250) PRIMARY KEY,
    email varchar(250)
)`
    USERS_INSERT="INSERT INTO users (name, email) VALUES(?,?)"
)

func create_table(db *sql.DB) {
    ctx := context.Background()
    _, err := db.ExecContext(ctx, USERS_TABLE)
    if err != nil {
        panic(err)
    }
}

func insert_rows(db *sql.DB) {
    ctx := context.Background()
    result, err := db.ExecContext(ctx, USERS_INSERT,
    "John", "john@gmail.com")
    if err != nil {
        panic(err)
    }

    lastUserId, err := result.LastInsertId()
    if err != nil {
        panic(err)
    }
    numRows, err := result.RowsAffected()
    if err != nil {
        panic(err)
    }
    fmt.Printf("User ID: %d, Rows: %d\n", lastUserId, numRows)
}
