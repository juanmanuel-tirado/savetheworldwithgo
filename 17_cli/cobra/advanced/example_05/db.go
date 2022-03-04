package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/juanmanuel-tirado/savetheworldwithgo/17_cli/cobra/advanced/example_05/cmd"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:  "db",
	Long: "Root command",
}

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get user data",
	Args:  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Get user %s!!!\n", args[0])
	},
	ValidArgsFunction: UserGet,
}

func UserGet(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	rand.Seed(time.Now().UnixNano())
	if rand.Int()%2 == 0 {
		return []string{"John", "Mary"}, cobra.ShellCompDirectiveNoFileComp
	}
	return []string{"Ernest", "Rick", "Mary"}, cobra.ShellCompDirectiveNoFileComp
}

func init() {
	RootCmd.AddCommand(GetCmd, cmd.CompletionCmd)
}

func main() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
