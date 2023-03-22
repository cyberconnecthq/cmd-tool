package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var viperCmd = &cobra.Command{
	Use:   "viper",
	Short: "Convert a viper config string to gitops format",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(ToSimpleAlfredResult(toGitopsFormat(input)))
	},
}

var input string

func init() {
	rootCmd.AddCommand(viperCmd)
	viperCmd.Flags().StringVarP(&input, "input", "i", "", "Required.")
	if err := viperCmd.MarkFlagRequired("input"); err != nil {
		panic(err)
	}
	if err := viper.BindPFlag("input", viperCmd.Flags().Lookup("input")); err != nil {
		panic(err)
	}
}

func toGitopsFormat(input string) string {
	upper := strings.ToUpper(input)
	return strings.ReplaceAll(upper, ".", "_")
}
