package main

import (
"encoding/json"
"fmt"
)

func main() {

    aNumber, _ := json.Marshal(42)

    var recoveredNumber int = -1
    err := json.Unmarshal(aNumber, &recoveredNumber)
    if err!= nil {
        panic(err)
    }
    fmt.Println(recoveredNumber)


    aMap, _ := json.Marshal(map[string]int{"one":1,"two":2})

    recoveredMap := make(map[string]int)
    err = json.Unmarshal(aMap, &recoveredMap)
    if err != nil {
        panic(err)
    }
    fmt.Println(recoveredMap)
}
