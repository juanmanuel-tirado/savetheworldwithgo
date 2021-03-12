package main

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/spf13/cobra/doc"
    "os"
)

var RootCmd = &cobra.Command{
    Use: "test",
    Short: "Documented test",
    Long: "How to document a command",
    Example: "./main test",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Save the world with Go!!!")
    },
}

func main() {

    RootCmd.Flags().Bool("flag",true,"Some flag")

    header := &doc.GenManHeader{
        Title: "Test",
        Manual: "MyManual",
        Section: "1",
    }
    err := doc.GenManTree(RootCmd, header, ".")
    if err != nil {
        panic(err)
    }
    err = doc.GenMarkdownTree(RootCmd, ".")
    if err != nil {
        panic(err)
    }
    err = doc.GenReSTTree(RootCmd, ".")
    if err != nil {
        panic(err)
    }
    err = doc.GenYamlTree(RootCmd, ".")
    if err != nil {
        panic(err)
    }
    if err := RootCmd.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}
