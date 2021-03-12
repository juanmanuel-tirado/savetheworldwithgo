package main

import (
    "fmt"
    "github.com/spf13/cobra"
    "os"
)

var RootCmd = &cobra.Command{
    Use: "main",
    Short: "short message",
    Run: func(cmd *cobra.Command, args []string) {
        for i:=0;i<cnfg.Rep;i++ {
            fmt.Printf("[[--%s--]]\n", cnfg.Msg)
        }
    },
}

type Config struct {
    Msg string
    Rep int
}
var cnfg Config = Config{}

func init() {
    RootCmd.Flags().StringVar(&cnfg.Msg, "msg","Save the world with Go!!!","Message to show")
    RootCmd.Flags().IntVar(&cnfg.Rep,"rep",1, "Number of times to show the message")
    RootCmd.MarkFlagRequired("rep")
}

func main() {
    if err := RootCmd.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}
