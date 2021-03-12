package example_01

import (
    "math/rand"
    "testing"
    "time"
)

func BuildGraph(vertices int, edges int) [][]int {
    graph := make([][]int, vertices)
    for i:=0;i<len(graph);i++{
        graph[i] = make([]int,0,1)
    }
    for i:=0;i<edges;i++{
        from := rand.Intn(vertices)
        to := rand.Intn(vertices)
        graph[from]=append(graph[from],to)
    }

    return graph
}

func BenchmarkGraph(b *testing.B) {
    rand.Seed(time.Now().UnixNano())
    for i:=0;i<b.N;i++ {
        BuildGraph(100,20000)
    }
}
