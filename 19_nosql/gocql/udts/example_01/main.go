package main

import (
    "context"
    "fmt"
    "github.com/gocql/gocql"
    "time"
)

const LAPTOP_TYPE = `CREATE TYPE example.Laptop (
sn int,
model text,
memory int)`

const USERS_TABLE =`CREATE TABLE example.users (
user_id int PRIMARY KEY,
name text,
email text,
Laptop frozen<Laptop>)`

const INSERT = `INSERT INTO example.users (
user_id, name, email, Laptop) VALUES (?,?,?,?)`

type Laptop struct {
    Sn int `cql:"sn"`
    Model string `cql:"model"`
    Memory int `cql:"memory"`
}

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
    err = session.Query(LAPTOP_TYPE).WithContext(ctx).Exec()
    if err != nil {
        panic(err)
    }
    err = session.Query(USERS_TABLE).WithContext(ctx).Exec()
    if err != nil {
        panic(err)
    }

    err = session.Query(INSERT,1,"John","john@gmail.com",&Laptop{100,"Lenovo",10}).Exec()
    if err != nil {
        panic(err)
    }

    var retrieved Laptop
    err = session.Query("select laptop from users where user_id=1").Scan(&retrieved)
    fmt.Println("Retrieved", retrieved)
}
