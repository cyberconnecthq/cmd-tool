package cmd

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var checksumCmd = &cobra.Command{
	Use:   "checksum",
	Short: "Convert a evm address to checksum address",
	Run: func(cmd *cobra.Command, args []string) {
		checksum, err := toChecksumAddress(address)
		if err != nil {
			fmt.Println(ToSimpleAlfredResult("wrong address"))
		} else {
			fmt.Println(ToSimpleAlfredResult(checksum))
		}
	},
}

var address string

func init() {
	rootCmd.AddCommand(checksumCmd)
	checksumCmd.Flags().StringVarP(&address, "address", "a", "", "Required.")
	if err := checksumCmd.MarkFlagRequired("address"); err != nil {
		panic(err)
	}
	if err := viper.BindPFlag("address", checksumCmd.Flags().Lookup("address")); err != nil {
		panic(err)
	}
}

func toChecksumAddress(address string) (string, error) {
	addr, err := common.NewMixedcaseAddressFromString(address)
	if err != nil {
		return "", fmt.Errorf("toChecksumAddress %v: %w", address, err)
	}
	return addr.Address().String(), nil
}
