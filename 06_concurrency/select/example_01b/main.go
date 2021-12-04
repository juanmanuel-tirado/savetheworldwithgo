package main

import (
    "fmt"
    "strconv"
    "time"
)

func sendNumbers(out chan int) {
    for i:=0; i < 5; i++ {
        time.Sleep(time.Millisecond * 500)
        out <- i
    }
    fmt.Println("no more numbers")
    close(out)
}

func sendMsgs(out chan string) {
    for i:=0; i < 5; i++ {
        time.Sleep(time.Millisecond * 300)
        out <- strconv.Itoa(i)
    }
    fmt.Println("no more msgs")
    close(out)
}

func main() {
    numbers := make(chan int)
    msgs := make(chan string)

    go sendNumbers(numbers)
    go sendMsgs(msgs)

    closedNums, closedMsgs := false, false

    for !closedNums || !closedMsgs {
        select {
        case num, ok := <- numbers:
            if ok {
                fmt.Printf("number %d\n", num)
            } else {
                closedNums = true
            }
        case msg, ok := <- msgs:
            if ok {
                fmt.Printf("msg %s\n", msg)
            } else {
                closedMsgs = true
            }
        }
    }
}
