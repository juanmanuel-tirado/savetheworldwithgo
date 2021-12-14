package main

import (
    "fmt"
    "github.com/spf13/cobra"
    "os"
)

var RootCmd = &cobra.Command{
    Use: "hello",
    Short: "short message",
    Long: "Long message",
    Version: "v0.1.0",
    Example: "this is an example",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Save the world with Go!!!")
    },
}

func main() {
    if err := RootCmd.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}
