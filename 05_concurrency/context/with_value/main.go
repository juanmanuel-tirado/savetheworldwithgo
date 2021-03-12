package main

import (
    "context"
    "errors"
    "fmt"
)

func main() {

    f := func(ctx context.Context, a int, b int) (int,error) {

        switch ctx.Value("action") {
        case "+":
            return a + b,nil
        case "-":
            return a - b,nil
        default:
            return 0, errors.New("unknown action")
        }
    }

    ctx := context.WithValue(context.Background(), "action", "+")
    r, err := f(ctx,22,20)
    fmt.Println(r,err)
    ctx2 := context.WithValue(context.Background(), "action", "-")
    r, err = f(ctx2,22,20)
    fmt.Println(r,err)
}
