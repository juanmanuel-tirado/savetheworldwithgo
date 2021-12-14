package main

import (
    "errors"
    "fmt"
    "io"
)

type MyReader struct {
    data string
    from int
}

func(r *MyReader) Read(p []byte) (int, error) {
    if p == nil {
        return -1, errors.New("nil target array")
    }
    if len(r.data) <= 0 || r.from == len(r.data){
        return 0, io.EOF
    }
    n := len(r.data) - r.from
    if len(p) < n {
        n = len(p)
    }
    for i:=0;i < n; i++ {
        b := byte(r.data[r.from])
        p[i] = b
        r.from++
    }
    if r.from == len(r.data) {
        return n, io.EOF
    }
    return n, nil
}

func main() {
    target := make([]byte,5)
    empty := MyReader{}
    n, err := empty.Read(target)
    fmt.Printf("Read %d: Error: %v\n",n,err)
    mr := MyReader{"Save the world with Go!!!",0}
    n, err = mr.Read(target)
    for err == nil {
        fmt.Printf("Read %d: Error: %v -> %s\n",n,err, target)
        n, err = mr.Read(target)
    }
    fmt.Printf("Read %d: Error: %v -> %s\n",n,err, target)
}
