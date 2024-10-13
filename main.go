package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"multitenant-hosting-cli/cmd"
)

var rootCmd = &cobra.Command{
	Use:   "multitenant-hosting-cli",
	Short: "A command-line interface for multitenant-hosting platform",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to the cli for multitenant hosting platform!")
	},
}

func main() {
	rootCmd.AddCommand(cmd.CreateAppCmd)
	rootCmd.Execute()
}
