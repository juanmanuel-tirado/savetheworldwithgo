package main

import (
	"context"
    "github.com/gocql/gocql"
    "time"
)

const CREATE_TABLE=`CREATE TABLE example.user (
id int PRIMARY KEY,
name text,
email text
)
`

func main() {
    cluster := gocql.NewCluster("127.0.0.1:9042")
    cluster.Keyspace = "example"
    cluster.Consistency = gocql.Quorum
    cluster.Timeout = time.Minute
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
}
