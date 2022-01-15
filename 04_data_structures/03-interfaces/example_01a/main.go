package main

import "fmt"

type Greeter interface {
    SayHello() string
}

type Person struct{
    name string
}

func (p *Person) SayHello() string {
    return "Hi! This is me "+ p.name
}

func main() {

    var g Greeter

    p := Person{"John"}
    // g = p --> Does not work
    g = &p
    fmt.Println(g.SayHello())
}
