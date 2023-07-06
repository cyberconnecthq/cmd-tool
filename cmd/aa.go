package cmd

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var abstractAccountCmd = &cobra.Command{
	Use:   "abstract-account",
	Short: "Calculates the abstract account address of a given address",
	Run: func(cmd *cobra.Command, args []string) {
		aaAddress, err := calculateAAAddress(eoa)
		if err != nil {
			fmt.Println(ToSimpleAlfredResult("wrong address"))
		} else {
			fmt.Println(ToSimpleAlfredResult(aaAddress))
		}
	},
}

var eoa string

func init() {
	rootCmd.AddCommand(abstractAccountCmd)
	abstractAccountCmd.Flags().StringVarP(&eoa, "eoa", "a", "", "Required.")
	if err := abstractAccountCmd.MarkFlagRequired("eoa"); err != nil {
		panic(err)
	}
	if err := viper.BindPFlag("eoa", abstractAccountCmd.Flags().Lookup("eoa")); err != nil {
		panic(err)
	}
}

func calculateAAAddress(eoa string) (string, error) {
	saltBytes := []byte{}
	validator := "0xf94e5a47150d20c4b804c30b6699d786549a5821"
	kernelTemplate := "0x27D57664e1CC984595F2F9234d6553979831167e"
	nextTemplate := "0xddE1C02fC27B7BE195227e1C5e13DfC15280366F"
	factory := "0x2DAB5E3e3449b5CaDf5126154fAbFe6d1e0e8aaD"
	data := eoa
	creationCode := common.Hex2Bytes("608060405260405161034a38038061034a833981016040819052610022916101ca565b6001600160a01b0382166100965760405162461bcd60e51b815260206004820152603060248201527f4549503139363750726f78793a20696d706c656d656e746174696f6e2069732060448201526f746865207a65726f206164647265737360801b60648201526084015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc82815581511561017c576000836001600160a01b0316836040516100db9190610298565b600060405180830381855af49150503d8060008114610116576040519150601f19603f3d011682016040523d82523d6000602084013e61011b565b606091505b505090508061017a5760405162461bcd60e51b815260206004820152602560248201527f4549503139363750726f78793a20636f6e7374727563746f722063616c6c2066604482015264185a5b195960da1b606482015260840161008d565b505b5050506102b4565b634e487b7160e01b600052604160045260246000fd5b60005b838110156101b557818101518382015260200161019d565b838111156101c4576000848401525b50505050565b600080604083850312156101dd57600080fd5b82516001600160a01b03811681146101f457600080fd5b60208401519092506001600160401b038082111561021157600080fd5b818501915085601f83011261022557600080fd5b81518181111561023757610237610184565b604051601f8201601f19908116603f0116810190838211818310171561025f5761025f610184565b8160405282815288602084870101111561027857600080fd5b61028983602083016020880161019a565b80955050505050509250929050565b600082516102aa81846020870161019a565b9190910192915050565b6088806102c26000396000f3fe60806040526000602d7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5490565b90503660008037600080366000845af43d6000803e808015604d573d6000f35b3d6000fdfea264697066735822122096a7204edc57708c92df34dceec1545100ba4a90e1185cca3261cdac2dc43ed864736f6c634300080e0033")
	saltBytes = append(saltBytes, common.HexToAddress(validator).Bytes()...)
	saltBytes = append(saltBytes, common.HexToAddress(data).Bytes()...)
	saltBytes = append(saltBytes, common.Hex2Bytes("0000000000000000000000000000000000000000000000000000000000000000")...)
	salt := crypto.Keccak256(saltBytes)

	encodeCallBytesPrefix := common.Hex2Bytes("cf7a1d77")

	addrType, err := abi.NewType("address", "", []abi.ArgumentMarshaling{})
	if err != nil {
		return "", err
	}
	bytesType, err := abi.NewType("bytes", "", []abi.ArgumentMarshaling{})
	if err != nil {
		return "", err
	}
	arguments := abi.Arguments{
		{
			Type: addrType,
		},
		{
			Type: addrType,
		},
		{
			Type: bytesType,
		},
	}

	encodeCallBytesPostfix, err := arguments.Pack(
		common.HexToAddress(validator),
		common.HexToAddress(nextTemplate),
		common.HexToAddress(data).Bytes(),
	)

	encodeCallBytes := append(encodeCallBytesPrefix, encodeCallBytesPostfix...)

	encodeBytes, _ := abi.Arguments{
		{
			Type: addrType,
		},
		{
			Type: bytesType,
		},
	}.Pack(common.HexToAddress(kernelTemplate), encodeCallBytes)

	encodePackedBytes := append(creationCode, encodeBytes...)

	salt32 := [32]byte{}
	for i := 0; i < len(salt32); i++ {
		salt32[i] = salt[i]
	}

	return crypto.CreateAddress2(common.HexToAddress(factory), salt32, crypto.Keccak256(encodePackedBytes)).Hex(), nil
}
