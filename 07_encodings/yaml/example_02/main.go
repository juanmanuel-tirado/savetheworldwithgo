package main

import (
    "fmt"
    "gopkg.in/yaml.v2"
)

func main() {

    aNumber, _ := yaml.Marshal(42)

    var recoveredNumber int = -1
    err := yaml.Unmarshal(aNumber, &recoveredNumber)
    if err!= nil {
        panic(err)
    }
    fmt.Println(recoveredNumber)


    aMap, _ := yaml.Marshal(map[string]int{"one":1,"two":2})

    recoveredMap := make(map[string]int)
    err = yaml.Unmarshal(aMap, &recoveredMap)
    if err != nil {
        panic(err)
    }
    fmt.Println(recoveredMap)
}
