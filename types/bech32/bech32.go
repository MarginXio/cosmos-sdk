package bech32

import (
	"errors"
	"fmt"
	"github.com/cosmos/btcutil/bech32"
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

// ConvertAndEncode converts from a base64 encoded byte string to base32 encoded byte string and then to bech32.
func RealConvertAndEncode(hrp string, data []byte) (string, error) {
	converted, err := bech32.ConvertBits(data, 8, 5, true)
	if err != nil {
		return "", fmt.Errorf("encoding bech32 failed: %w", err)
	}

	return bech32.Encode(hrp, converted)
}

// DecodeAndConvert decodes a bech32 encoded string and converts to base64 encoded bytes.
func RealDecodeAndConvert(bech string) (string, []byte, error) {
	hrp, data, err := bech32.Decode(bech, 1023)
	if err != nil {
		return "", nil, fmt.Errorf("decoding bech32 failed: %w", err)
	}

	converted, err := bech32.ConvertBits(data, 5, 8, false)
	if err != nil {
		return "", nil, fmt.Errorf("decoding bech32 failed: %w", err)
	}

	return hrp, converted, nil
}
