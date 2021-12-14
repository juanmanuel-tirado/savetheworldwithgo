package main

import "log"

func main() {
    // Panic or fatal messages stop the execution flow
    // log.Fatal("This is a fatal message")
    // log.Panic("This is a panic message")
    log.Println("This is a log message")
    log.SetPrefix("prefix -> ")
    log.Println("This is a log message")
    log.SetFlags(log.Lshortfile)

}
