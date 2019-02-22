package cmd

import (
	"fmt"
	"log"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var inputFile string
var nameServer string
var nameServerPort string
var recordType string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "drl",
	Short: "DNS Record Lookup",
	Long: `drl: DNS Record Lookup

drl is a utility for looking up DNS records. It features a simple command
syntax, and easy-to-read results.

Complete documentation is at https://github.com/cconover/drl`,
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 && inputFile == "" {
			cmd.Help()
			os.Exit(0)
		}

		// A record lookup
		a := A{}
		for _, name := range args {
			resp, err := a.Get(name)
			if err != nil {
				log.Panic(err)
			}
			for _, record := range resp {
				fmt.Printf("%s %s\n", name, record)
				fmt.Println(name, record)
			}
		}
		os.Exit(0)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.drl.yaml)")
	rootCmd.PersistentFlags().StringVarP(&inputFile, "input-file", "i", "", "Text file containing a list of names for lookup, each on a new line.")
	rootCmd.PersistentFlags().StringVar(&nameServer, "nameserver", "8.8.8.8", "Specify the name server to query.")
	rootCmd.PersistentFlags().StringVar(&nameServerPort, "nameserver-port", "53", "Specify the port over which to connect to the name server.")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".drl" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".drl")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
