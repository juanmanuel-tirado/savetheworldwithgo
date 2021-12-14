package main

import (
    "encoding/csv"
    "fmt"
    "io"
    "log"
    "strings"
)


func main() {
    in := `user_id,score,password
"Gopher",1000,"admin"
"BigJ",10,"1234"
"GGBoom",,"1111"
`
    r := csv.NewReader(strings.NewReader(in))

    for {
        record, err := r.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(record)
    }
}
