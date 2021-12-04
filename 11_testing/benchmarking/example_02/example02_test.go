package example_02

import "testing"

func Sum(n int64) int64 {
    var result int64 = 0
    var i int64
    for i = 0; i<n; i++ {
        result = result + i
    }
    return result
}

func BenchmarkSumParallel(b *testing.B) {
    b.RunParallel(func(pb *testing.PB){
        for pb.Next() {
            Sum(1000000)
        }
    })
}

