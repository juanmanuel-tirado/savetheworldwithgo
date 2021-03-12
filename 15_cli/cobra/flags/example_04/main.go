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
        for i:=0;i<*Rep;i++ {
            fmt.Printf("[[--%s--]]\n", *Msg)
        }
    },
}

var Msg *string
var Rep *int

func init() {
    Msg = RootCmd.Flags().String("msg", "Save the world with Go!!!",
        "Message to show")
    Rep = RootCmd.Flags().Int("rep",1, "Number of times to show the message")

    RootCmd.MarkFlagRequired("rep")
}

func main() {
    if err := RootCmd.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}
