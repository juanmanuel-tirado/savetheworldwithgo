package main

import "fmt"

func main(){
    sales := map[string]int {
        "Jan": 34345,
        "Feb": 11823,
        "Mar": 8838,
        "Apr": 33,
    }

    fmt.Println("Month\tSales")
    for month, sale := range sales {
        fmt.Printf("%s\t\t%d\n",month,sale)
    }
}
