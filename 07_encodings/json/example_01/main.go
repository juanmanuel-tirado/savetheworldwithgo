package main

import (
"encoding/json"
"fmt"
)

func main() {
    number, err := json.Marshal(42)
    if err!=nil{
        panic(err)
    }
    fmt.Println(string(number))

    float, _ := json.Marshal(3.14)
    fmt.Println(string(float))

    msg, _ := json.Marshal("This is a msg!!!")
    fmt.Println(string(msg))

    numbers, _ := json.Marshal([]int{1,1,2,3,5,8})
    fmt.Println(string(numbers))

    aMap, _ := json.Marshal(map[string]int{"one":1,"two":2})
    fmt.Println(string(aMap))
}
