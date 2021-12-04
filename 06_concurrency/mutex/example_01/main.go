package main

import (
    "fmt"
    "sync"
    "time"
)

func writer(x map[int]int, factor int, m *sync.Mutex) {
    i := 1
    for {
        time.Sleep(time.Second)
        m.Lock()
        x[i] = x[i-1] * factor
        m.Unlock()
        i++
    }
}

func reader(x map[int]int, m *sync.Mutex) {
    for {
        time.Sleep(time.Millisecond*500)
        m.Lock()
        fmt.Println(x)
        m.Unlock()
    }
}

func main() {
    x := make(map[int]int)
    x[0]=1

    m := sync.Mutex{}
    go writer(x, 2, &m)
    go reader(x, &m)

    time.Sleep(time.Millisecond * 300)
    go writer(x, 3, &m)

    time.Sleep(time.Second*4)
}
