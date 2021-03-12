package main

import "fmt"

func main() {
    birthdays := map[string]string{
        "Jesus": "12-25-0000",
        "Budha": "563 BEC",
    }
    fmt.Println(birthdays,len(birthdays))

    xmas, found := birthdays["Jesus"]
    fmt.Println(xmas, found)

    delete(birthdays, "Jesus")
    fmt.Println(birthdays,len(birthdays))

    _, found = birthdays["Jesus"]
    fmt.Println("Did we find when its Xmas?", found)

    birthdays["Jesus"]="12-25-0000"
    fmt.Println(birthdays)

}
