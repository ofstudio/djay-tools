package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	SilenceUsage:      true,
	Use:               "djay-tools",
	Short:             "Command-line companion for Algoriddim djay Pro.",
	Long:              `This tool provides various utilities for Algoriddim djay Pro.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
