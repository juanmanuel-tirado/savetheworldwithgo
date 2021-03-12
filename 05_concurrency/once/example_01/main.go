package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

var first int

func setter(i int, ch chan bool, once *sync.Once) {
    t := rand.Uint32() % 300
    time.Sleep(time.Duration(t)*time.Millisecond)
    once.Do(func(){
        first = i
    })
    ch <- true
    fmt.Println(i,"Done")
}

func main() {
    rand.Seed(time.Now().UnixNano())

    var once sync.Once

    ch := make(chan bool)
    for i:=0;i<5;i++ {
       go setter(i, ch, &once)
    }

    for i:=0;i<5;i++{
        <- ch
    }
    fmt.Println("The first was", first)
}
