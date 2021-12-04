package main

import (
    "fmt"
    "github.com/spf13/cobra"
    "os"
)

var RootCmd = &cobra.Command{
    Use: "say",
    Long: "Root command",
    PersistentPreRun: func(cmd *cobra.Command, args []string) {
        fmt.Printf("Hello %s!!!\n", person)
    },
    Run: func(cmd *cobra.Command, args []string) {},
    PostRun: func(cmd *cobra.Command, args []string) {
        fmt.Printf("Bye %s!!!\n", person)
    },
}

var SomethingCmd = &cobra.Command{
    Use: "something",
    Short: "Say something",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Printf("%s\n",msg)
    },
    PostRun: func(cmd *cobra.Command, args []string) {
        fmt.Printf("That's all I have to say %s\n", person)
    },
}

var person string
var msg string

func init() {
    RootCmd.AddCommand(SomethingCmd)
    RootCmd.Flags().StringVar(&person, "person", "Mr X", "Receiver")
    SomethingCmd.Flags().StringVar(&msg, "msg", "", "Message to say")
    SomethingCmd.MarkFlagRequired("msg")
}

func main() {
    if err := RootCmd.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}
