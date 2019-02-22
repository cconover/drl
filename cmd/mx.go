package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// mxCmd represents the mx command
var mxCmd = &cobra.Command{
	Use:   "mx",
	Short: "MX record lookup",
	Long:  `Look up the MX record(s) for the provided name(s).`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mx called")
	},
}

func init() {
	rootCmd.AddCommand(mxCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mxCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mxCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
