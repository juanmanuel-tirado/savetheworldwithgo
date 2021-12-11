package main

import "fmt"

func main(){
    names := []string{
        "Jeremy", "John", "Joseph",
    }

    for i:=0;i<len(names);i++{
        fmt.Println(i,names[i])
    }

    for position, name := range(names){
        fmt.Println(position,name)
    }
}
