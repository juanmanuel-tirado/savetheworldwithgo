package main

import (
    "fmt"
    "github.com/spf13/cobra"
    "os"
)

var RootCmd = &cobra.Command{
    Use: "main",
    Long: "Long message",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Printf("[[--%s--]]\n", *Msg)
    },
}

var Msg *string

func init() {
    Msg = RootCmd.Flags().String("msg", "Save the world with Go!!!",
        "Message to show")
}

func main() {
    if err := RootCmd.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}
