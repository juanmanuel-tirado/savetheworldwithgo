package main

import (
    "fmt"
    "sync"
    "sync/atomic"
    "time"
)

type Monitor struct {
    ActiveUsers int
    Requests int
}

func updater(monitor atomic.Value, m *sync.Mutex) {
    for {
        time.Sleep(time.Millisecond*500)
        m.Lock()
        current := monitor.Load().(*Monitor)
        current.ActiveUsers += 100
        current.Requests += 300
        monitor.Store(current)
        m.Unlock()
    }
}

func observe(monitor atomic.Value) {
    for {
        time.Sleep(time.Second)
        current := monitor.Load()
        fmt.Printf("%v\n", current)
    }
}

func main() {
    var monitor atomic.Value
    monitor.Store(&Monitor{0,0})
    m := sync.Mutex{}

    go updater(monitor, &m)
    go observe(monitor)

    time.Sleep(time.Second * 5)
}
