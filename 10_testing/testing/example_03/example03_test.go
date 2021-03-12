package example_03

import (
    "fmt"
    "math"
    "testing"
    "time"
)

func Sum(n int64, done chan bool) {
    var result int64 = 0
    var i int64
    for i = 0; i<n; i++ {
        result = result + i
    }
    done <- true
}

func TestSum(t *testing.T) {
    var i int64
    done := make(chan bool)
    for i = 1000; i<math.MaxInt64; i+=100000 {
        go Sum(i, done)
        timeout := time.NewTimer(time.Millisecond)
        select {
        case <- timeout.C:
            t.Skip(fmt.Sprintf("%d took longer than 1 millisecond",i))
        case <- done:

        }
    }
}
