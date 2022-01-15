package main

import "fmt"

type Animal interface {
    Roar() string
    Run() string
}

type Dog struct {}

func (d Dog) Roar() string {
    return "woof"
}

func (d Dog) Run() string {
    return "run like a dog"
}

type Cat struct {}

func (c *Cat) Roar() string {
    return "meow"
}

func (c *Cat) Run() string {
    return "run like a cat"
}

func RoarAndRun(a Animal) {
    fmt.Printf("%s and %s\n", a.Roar(), a.Run())
}

func main() {
    myDog := Dog{}
    myCat := Cat{}

    RoarAndRun(myDog)
    RoarAndRun(&myCat)
}
