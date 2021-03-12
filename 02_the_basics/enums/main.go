package main

import "fmt"

type DayOfTheWeek uint8

const(
    Monday DayOfTheWeek = iota
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
    Sunday
)

func main() {

    fmt.Printf("Monday is %d\n", Monday)
    fmt.Printf("Wednesday is %d\n", Wednesday)
    fmt.Printf("Friday is %d\n", Friday)
}
