package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

func generator(ch chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    for i:=0;i<5;i++ {
        time.Sleep(time.Millisecond*200)
        ch <- rand.Int()
    }
    close(ch)
    fmt.Println("Generator done")
}

func consumer(id int, sleep time.Duration,
    ch chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    for task := range(ch) {
        time.Sleep(time.Millisecond*sleep)
        fmt.Printf("%d - task[%d]\n",id,task)
    }
    fmt.Printf("Consumer %d done\n",id)
}

func main() {
    rand.Seed(42)

    ch := make(chan int,10)
    var wg sync.WaitGroup
    wg.Add(3)

    go generator(ch,&wg)
    go consumer(1,400,ch,&wg)
    go consumer(2,100,ch,&wg)

    wg.Wait()
}
