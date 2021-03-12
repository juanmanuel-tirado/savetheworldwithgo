package main

import (
    "context"
    "github.com/gocql/gocql"
)

const CREATE_TABLE=`CREATE TABLE users (
id int PRIMARY KEY,
name text,
email text
)`

const INSERT_QUERY=`INSERT INTO users
(id,name,email) VALUES(?,?,?)`

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

    err = session.Query(INSERT_QUERY,1, "John", "john@gmail.com").WithContext(ctx).Exec()
    if err != nil {
        panic(err)
    }
}
