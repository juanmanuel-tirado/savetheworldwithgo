package example_01

import (
    "fmt"
    "testing"
)

func Sum(n int64) int64 {
    var result int64 = 0
    var i int64
    for i = 0; i<n; i++ {
        result = result + i
    }
    return result
}

func BenchmarkSum(b *testing.B) {
    fmt.Println("b.N:",b.N)
    for i:=0;i<b.N;i++ {
        Sum(1000000)
    }
}
