package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println(">>> What do you have to say?\n")
    counter := 0
    for scanner.Scan() {
        text := scanner.Text()
        counter = counter + len(text)
        if counter > 15 {
            break
        }
    }
    fmt.Println("that's enough")
}
