package cmd

import (
	"net"

	"github.com/spf13/cobra"
)

// The A struct contains the methods and properties relevant to working with
// A records.
type A struct {
	Name  string
	Value string
	TTL   string
}

// aCmd represents the a command
var aCmd = &cobra.Command{
	Use:   "a",
	Short: "A record lookup",
	Long:  `Look up the A record for the provided name(s).`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(aCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// aCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// aCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Get retrieves the DNS A record for a specified address.
func (a *A) Get(name string) ([]net.IP, error) {
	addr, err := net.LookupIP(name)
	if err != nil {
		return addr, err
	}

	return addr, nil
}
