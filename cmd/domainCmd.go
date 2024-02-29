package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var domainCmd = &cobra.Command{
	Use:   "domain",
	Short: "use domain for querying a domain1",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		domain := args[0]
		domainCommand(domain)
	},
}

func domainCommand(domain string) {
	log.Printf("analysing domain: %v", domain)
}
