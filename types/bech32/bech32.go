package bech32

import (
	"github.com/ethereum/go-ethereum/common"
)

const prefix = "0x"

func ConvertAndEncode(hrp string, data []byte) (string, error) {
	if len(data) == 0 {
		panic("empty bytes")
	}
	return common.BytesToAddress(data).String(), nil
}

// DecodeAndConvert decodes a bech32 encoded string and converts to base64 encoded bytes.
func DecodeAndConvert(bech string) (string, []byte, error) {
	return prefix, common.HexToAddress(bech).Bytes(), nil
}