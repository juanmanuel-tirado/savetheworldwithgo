package main

import "fmt"

func main(){
    names := []string{
        "Jeremy", "John", "Joseph",
    }

    for _, name := range(names){
        name = name + "_changed"
    }
    fmt.Println(names)

    for position, name := range(names){
        names[position] = name + "_changed"
    }
    fmt.Println(names)


}
