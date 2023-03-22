package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cmd-tool",
	Short: "CyberConnect cmd tools",
	Long:  "A cmd tools Go implementation for CyberConnect.",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}
