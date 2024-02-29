package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"net/netip"
)

var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "use ip for querying an IP address",
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(1)(cmd, args); err != nil {
			return err
		}

		if _, err := netip.ParseAddr(args[0]); err != nil {
			return fmt.Errorf("not an IP address: %s", args[0])
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		addr, err := netip.ParseAddr(args[0])
		if err != nil {
			log.Default().Fatalf("can't invoke ip cmd. not an IP address: (%s). should have been catched in validation.", args[0])
		}
		ipCommand(addr)
	},
}

func ipCommand(addr netip.Addr) {
	ipVersion := getIPVersion(addr)
	log.Printf("analysing IPv%s IP: %v", ipVersion, addr)
}

func getIPVersion(addr netip.Addr) string {
	var ipVersion = "UNKNOWN"
	if addr.Is4() {
		ipVersion = "4"
	} else if addr.Is6() {
		ipVersion = "6"
	}
	return ipVersion
}
