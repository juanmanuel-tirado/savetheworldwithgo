package main

import (
    "context"
    "fmt"
    "github.com/gocql/gocql"
)

const (
	QUERY        ="SELECT * FROM users"
    CREATE_TABLE =`CREATE TABLE users (
id int PRIMARY KEY,
name text,
email text
)`
    INSERT_QUERY =`INSERT INTO users
(id,name,email) VALUES(?,?,?)`
)

func main() {
    cluster := gocql.NewCluster("127.0.0.1:9042")
    cluster.Keyspace = "example"
    session, err := cluster.CreateSession()
    if err != nil {
        panic(err)
    }
    defer session.Close()

    ctx := context.Background()
    err = session.Query(CREATE_TABLE).WithContext(ctx).Exec()
    if err != nil {
        panic(err)
    }

    err = session.Query(INSERT_QUERY,2, "Mary", "mary@gmail.com").WithContext(ctx).Exec()
    if err != nil {
        panic(err)
    }
    err = session.Query(INSERT_QUERY,1, "John", "john@gmail.com").WithContext(ctx).Exec()
    if err != nil {
        panic(err)
    }

    scanner := session.Query(QUERY).WithContext(ctx).Iter().Scanner()
    for scanner.Next() {
        var id int
        var  name, email string
        err = scanner.Scan(&id,&name,&email)
        if err != nil {
            panic(err)
        }
        fmt.Println("Found:",id,name,email)
    }
}
