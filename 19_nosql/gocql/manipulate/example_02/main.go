package main

import (
    "context"
    "github.com/gocql/gocql"
)

const CREATE_TABLE=`CREATE TABLE users (
id int PRIMARY KEY,
name text,
email text)`

const INSERT_QUERY=`INSERT INTO users
(id,name,email) VALUES(?,?,?)`

const UPDATE_QUERY=`UPDATE users SET email=? WHERE id=?`

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

    b := session.NewBatch(gocql.UnloggedBatch).WithContext(ctx)
    b.Entries = append(b.Entries,
        gocql.BatchEntry {
            Stmt: INSERT_QUERY,
            Args: []interface{}{1, "John", "john@gmail.com"},
        },
        gocql.BatchEntry {
            Stmt: UPDATE_QUERY,
            Args: []interface{}{"otheremail@email.com",1},
        })
    err = session.ExecuteBatch(b)
    if err != nil {
        panic(err)
    }
}
