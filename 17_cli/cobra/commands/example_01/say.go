package main

import (
    "fmt"
    "github.com/spf13/cobra"
    "os"
)

var RootCmd = &cobra.Command{
    Use: "say",
    Long: "Root command",
}

var HelloCmd = &cobra.Command{
    Use: "hello",
    Short: "Say hello",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Hello!!!")
    },
}

var ByeCmd = &cobra.Command{
    Use: "bye",
    Short: "Say goodbye",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Bye!!!")
    },
}

func init() {
    RootCmd.AddCommand(HelloCmd, ByeCmd)
}

func main() {
    if err := RootCmd.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}
