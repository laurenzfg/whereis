package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "whereis",
	Short: "Whereis is a tool to locate oneself and other resources on the Internet",
}

func init() {
	rootCmd.AddCommand(ipCmd)
	rootCmd.AddCommand(domainCmd)
}

func Execute() int {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		return 1
	}
	return 0
}
