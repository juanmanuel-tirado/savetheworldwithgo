package main

import (
    "os"
)

func main() {
    filePath := "/tmp/msg"
    msg := []string{
        "Rule", "the", "world", "with", "Go!!!"}

    f, err := os.Create(filePath)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    for _, s := range msg {
        f.WriteString(s+"\n")
    }
}
