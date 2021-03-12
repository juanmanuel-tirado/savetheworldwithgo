package main

import (
    "encoding/csv"
    "os"
)

func main() {
    out := [][]string{
        {"user_id","score","password"},
        {"Gopher","1000","admin"},
        {"BigJ","10","1234"},
        {"GGBoom","","1111"},
    }
    writer := csv.NewWriter(os.Stdout)
    for _, rec := range out {
        err := writer.Write(rec)
        if err != nil {
            panic(err)
        }
    }
    writer.Flush()

}

