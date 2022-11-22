package ledger

import (
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ethsecp256k1"
	"github.com/cosmos/cosmos-sdk/crypto/types"
)

type (
	PrivKeyLedgerETHSecp256k1 struct {
		CachedPubKey types.PubKey
		Path         hd.BIP44Params
	}
)

func NewPrivKeyETHSecp256k1Unsafe(path hd.BIP44Params) (types.LedgerPrivKey, error) {
	device, err := getDevice()
	if err != nil {
		return nil, err
	}
	defer warnIfErrors(device.Close)

	pubKey, err := getETHPubKeyUnsafe(device, path)
	if err != nil {
		return nil, err
	}

	return PrivKeyLedgerETHSecp256k1{pubKey, path}, nil
}

func NewPrivKeyETHSecp256k1(path hd.BIP44Params, hrp string) (types.LedgerPrivKey, string, error) {
	device, err := getDevice()
	if err != nil {
		return nil, "", fmt.Errorf("failed to retrieve device: %w", err)
	}
	defer warnIfErrors(device.Close)

	pubKey, addr, err := getETHPubKeyAddrSafe(device, path, hrp)
	if err != nil {
		return nil, "", fmt.Errorf("failed to recover pubkey: %w", err)
	}

	return PrivKeyLedgerETHSecp256k1{pubKey, path}, addr, nil
}

// PubKey returns the cached public key.
func (pkl PrivKeyLedgerETHSecp256k1) PubKey() types.PubKey {
	return pkl.CachedPubKey
}

// Sign returns a secp256k1 signature for the corresponding message
func (pkl PrivKeyLedgerETHSecp256k1) Sign(message []byte) ([]byte, error) {
	device, err := getDevice()
	if err != nil {
		return nil, err
	}
	defer warnIfErrors(device.Close)

	return ethSign(device, pkl, message)
}

func ethSign(device SECP256K1, pkl PrivKeyLedgerETHSecp256k1, msg []byte) ([]byte, error) {
	err := validateETHKey(device, pkl)
	if err != nil {
		return nil, err
	}

	sig, err := device.SignSECP256K1(pkl.Path.DerivationPath(), msg)
	if err != nil {
		return nil, err
	}

	return convertDERtoBER(sig)
}

// ValidateKey allows us to verify the sanity of a public key after loading it
// from disk.
func (pkl PrivKeyLedgerETHSecp256k1) ValidateKey() error {
	device, err := getDevice()
	if err != nil {
		return err
	}
	defer warnIfErrors(device.Close)

	return validateETHKey(device, pkl)
}

func validateETHKey(device SECP256K1, pkl PrivKeyLedgerETHSecp256k1) error {
	pub, err := getETHPubKeyUnsafe(device, pkl.Path)
	if err != nil {
		return err
	}

	// verify this matches cached address
	if !pub.Equals(pkl.CachedPubKey) {
		return fmt.Errorf("cached key does not match retrieved key")
	}

	return nil
}

func getETHPubKeyUnsafe(device SECP256K1, path hd.BIP44Params) (types.PubKey, error) {
	publicKey, err := device.GetPublicKeySECP256K1(path.DerivationPath())
	if err != nil {
		return nil, fmt.Errorf("please open Cosmos app on the Ledger device - error: %v", err)
	}

	// re-serialize in the 33-byte compressed format
	cmp, err := btcec.ParsePubKey(publicKey, btcec.S256())
	if err != nil {
		return nil, fmt.Errorf("error parsing public key: %v", err)
	}

	compressedPublicKey := make([]byte, ethsecp256k1.PubKeySize)
	copy(compressedPublicKey, cmp.SerializeCompressed())

	return &ethsecp256k1.PubKey{Key: compressedPublicKey}, nil
}

// AssertIsPrivKeyInner implements the PrivKey interface. It performs a no-op.
func (pkl *PrivKeyLedgerETHSecp256k1) AssertIsPrivKeyInner() {}

// Bytes implements the PrivKey interface. It stores the cached public key so
// we can verify the same key when we reconnect to a ledger.
func (pkl PrivKeyLedgerETHSecp256k1) Bytes() []byte {
	return cdc.MustMarshal(pkl)
}

// Equals implements the PrivKey interface. It makes sure two private keys
// refer to the same public key.
func (pkl PrivKeyLedgerETHSecp256k1) Equals(other types.LedgerPrivKey) bool {
	if otherKey, ok := other.(PrivKeyLedgerETHSecp256k1); ok {
		return pkl.CachedPubKey.Equals(otherKey.CachedPubKey)
	}
	return false
}

func (pkl PrivKeyLedgerETHSecp256k1) Type() string { return "PrivKeyLedgerETHSecp256k1" }

func getETHPubKeyAddrSafe(device SECP256K1, path hd.BIP44Params, hrp string) (types.PubKey, string, error) {
	publicKey, addr, err := device.GetAddressPubKeySECP256K1(path.DerivationPath(), hrp)
	if err != nil {
		return nil, "", fmt.Errorf("%w: address rejected for path %s", err, path.String())
	}

	cmp, err := btcec.ParsePubKey(publicKey, btcec.S256())
	if err != nil {
		return nil, "", fmt.Errorf("error parsing public key: %v", err)
	}

	compressedPublicKey := make([]byte, ethsecp256k1.PubKeySize)
	copy(compressedPublicKey, cmp.SerializeCompressed())

	return &ethsecp256k1.PubKey{Key: compressedPublicKey}, addr, nil
}