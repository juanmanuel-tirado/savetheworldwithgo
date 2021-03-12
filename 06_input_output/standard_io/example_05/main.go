package main

import (
    "bufio"
    "os"
    "time"
)

func main() {
    writer := bufio.NewWriter(os.Stdout)

    msg := "Save the world with Go!!!"
    for _, letter := range msg {
        time.Sleep(time.Millisecond*300)
        writer.WriteByte(byte(letter))
        writer.Flush()
    }
}
