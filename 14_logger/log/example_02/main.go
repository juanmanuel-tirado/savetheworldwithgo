package main

import (
    "io/ioutil"
    "log"
    "os"
)

func main() {
    tmpFile, err := ioutil.TempFile(os.TempDir(),"logger.out")
    if err != nil {
        log.Panic(err)
    }
    logger := log.New(tmpFile, "prefix -> ", log.Ldate)
    logger.Println("This is a log message")
}
