package main

import "fmt"

func main() {
    var ages map[string]int
    fmt.Println(ages)

    // This fails, ages was not initialized
    // ages["Jesus"] = 33

    ages = make(map[string]int,5)
    // Now it works because it was initialized
    ages["Jesus"] = 33

    ages = map[string]int{
        "Jesus": 33,
        "Mathusalem": 969,
    }
    fmt.Println(ages)
}
