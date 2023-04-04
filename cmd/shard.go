package cmd

import (
	"fmt"
	"hash/crc32"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var shardCmd = &cobra.Command{
	Use:   "shard",
	Short: "Calculate shard number from a sharding key",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(ToSimpleAlfredResult(calculateShardingNumber(shardingKey)))
	},
}

var shardingKey string

func init() {
	rootCmd.AddCommand(shardCmd)
	shardCmd.Flags().StringVarP(&shardingKey, "sharding key", "s", "", "Required.")
	if err := shardCmd.MarkFlagRequired("sharding key"); err != nil {
		panic(err)
	}
	if err := viper.BindPFlag("sharding key", shardCmd.Flags().Lookup("sharding key")); err != nil {
		panic(err)
	}
}

func calculateShardingNumber(shardingKey string) string {
	id := 0

	id, err := strconv.Atoi(shardingKey)
	if err != nil {
		id = int(crc32.ChecksumIEEE([]byte(shardingKey)))
	}

	return fmt.Sprintf("%02d", id%32)
}
