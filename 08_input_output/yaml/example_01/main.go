package main

import (
    "gopkg.in/yaml.v2"
    "fmt"
)

func main() {
    number, err := yaml.Marshal(42)
    if err!=nil{
        panic(err)
    }
    fmt.Println(string(number))

    float, _ := yaml.Marshal(3.14)
    fmt.Println(string(float))

    msg, _ := yaml.Marshal("This is a msg!!!")
    fmt.Println(string(msg))

    numbers, _ := yaml.Marshal([]int{1,1,2,3,5,8})
    fmt.Println(string(numbers))

    aMap, _ := yaml.Marshal(map[string]int{"one":1,"two":2})
    fmt.Println(string(aMap))
}
