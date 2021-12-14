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
        fmt.Printf("Hello %s!!!\n",person)
    },
}

var ByeCmd = &cobra.Command{
    Use: "bye",
    Short: "Say goodbye",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Printf("Bye %s!!!\n",person)
    },
}

var CustomCmd = &cobra.Command{
    Use: "custom",
    Short: "Custom greetings",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Printf("Say %s to %s\n",msg,person)
    },
}

var msg string
var person string

func init() {
    RootCmd.AddCommand(HelloCmd, ByeCmd, CustomCmd)

    RootCmd.PersistentFlags().StringVar(&person, "person", "Mr X", "Receiver")
    CustomCmd.Flags().StringVar(&msg,"msg","what's up","Custom message")
}

func main() {
    if err := RootCmd.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}
