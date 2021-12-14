package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print(">>> What do you have to say?\n")
    fmt.Print("<<< ")
    text, err := reader.ReadString('\n')
    if err != nil {
        panic(err)
    }
    fmt.Println(">>> You're right!!!")
    fmt.Println(strings.ToUpper(text))
}
