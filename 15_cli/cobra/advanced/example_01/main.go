package main

import (
    "errors"
    "fmt"
    "github.com/spf13/cobra"
    "os"
)

var RootCmd = &cobra.Command{
    Use: "main",
    Short: "short message",
}

var ActionCmd = &cobra.Command{
    Use: "action",
    Args: cobra.MinimumNArgs(2),
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Do something with ",args)
    },
}

func helper (cmd *cobra.Command, args []string) {
    fmt.Printf("You entered command %s\n", cmd.Name())
    fmt.Println("And that is all the help we have right now :)")
}

func usager (cmd *cobra.Command) error {
    fmt.Printf("You entered command %s\n", cmd.Name())
    fmt.Println("And you do not know how it works :)")
    return errors.New("Something went wrong :(")
}

func main() {
    RootCmd.AddCommand(ActionCmd)
    RootCmd.SetHelpFunc(helper)
    RootCmd.SetUsageFunc(usager)

    ActionCmd.Flags().Bool("now",false,"Do it now")

    if err := RootCmd.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}
