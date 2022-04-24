package bech32

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"strings"
)

const prefix = "0x"

func ConvertAndEncode(hrp string, data []byte) (string, error) {
	if len(data) == 0 {
		return "", errors.New("empty bytes")
	}
	return common.BytesToAddress(data).String(), nil
}

// DecodeAndConvert decodes a bech32 encoded string and converts to base64 encoded bytes.
func DecodeAndConvert(bech string) (string, []byte, error) {
	if len(bech) != 42 {
		return prefix, nil, errors.New(fmt.Sprintf("invalid address length, expected: %d, got: %d", 42, len(bech)))
	}
	if !strings.HasPrefix(bech, prefix) {
		return prefix, nil, errors.New(fmt.Sprintf("invalid address prefix, expected %s", prefix))
	}
	ethAddress := common.HexToAddress(bech)
	commonAddress := common.Address{}
	if ethAddress == commonAddress {
		return prefix, nil, errors.New(fmt.Sprintf("invalid address"))
	}
	return prefix, ethAddress.Bytes(), nil
}
