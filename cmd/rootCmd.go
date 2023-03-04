package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "whereis",
	Short: "Whereis is a tool to locate oneself and other resources on the Internet",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Jo Jo Jo Whats Up?")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
