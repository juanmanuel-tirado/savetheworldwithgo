package main

import (
    "fmt"
    "github.com/spf13/cobra"
    "os"
    "strings"
)

var RootCmd = &cobra.Command{
    Use: "main",
    Long: "Long message",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Printf("%s\n",strings.Join(args,","))
    },
}

func main() {
    if err := RootCmd.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}
