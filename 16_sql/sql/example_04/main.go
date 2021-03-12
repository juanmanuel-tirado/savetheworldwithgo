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

    runTransaction(db)
}


func runTransaction(db *sql.DB) {

    tx, err := db.BeginTx(context.Background(), nil)
    defer tx.Rollback()

    ctx := context.Background()
    _, err = tx.ExecContext(ctx, "INSERT INTO users(name, email) VALUES(?,?)",
        "Peter", "peter@email.com")
    if err != nil {
        panic(err)
    }
    err = tx.Commit()
    if err != nil {
        panic(err)
    }
}
