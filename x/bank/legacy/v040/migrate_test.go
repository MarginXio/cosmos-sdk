package v040_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	v038auth "github.com/cosmos/cosmos-sdk/x/auth/legacy/v038"
	v039auth "github.com/cosmos/cosmos-sdk/x/auth/legacy/v039"
	v036supply "github.com/cosmos/cosmos-sdk/x/bank/legacy/v036"
	v038bank "github.com/cosmos/cosmos-sdk/x/bank/legacy/v038"
	v040bank "github.com/cosmos/cosmos-sdk/x/bank/legacy/v040"
)

func TestMigrate(t *testing.T) {
	encodingConfig := simapp.MakeTestEncodingConfig()
	clientCtx := client.Context{}.
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithJSONCodec(encodingConfig.Marshaler)

	coins := sdk.NewCoins(sdk.NewInt64Coin("stake", 50))
	addr1, _ := sdk.AccAddressFromBech32("0x31ADcCDBfdf9599930Bc983877b8ac0AC990a7a5")
	acc1 := v038auth.NewBaseAccount(addr1, coins, nil, 1, 0)

	addr2, _ := sdk.AccAddressFromBech32("0xA328f26C3A9d1b3F4723B017b03F87EB43f3b63c")
	vaac := v038auth.NewContinuousVestingAccountRaw(
		v038auth.NewBaseVestingAccount(
			v038auth.NewBaseAccount(addr2, coins, nil, 1, 0), coins, nil, nil, 3160620846,
		),
		1580309972,
	)

	supply := sdk.NewCoins(sdk.NewInt64Coin("stake", 1000))

	bankGenState := v038bank.GenesisState{
		SendEnabled: true,
	}
	authGenState := v039auth.GenesisState{
		Accounts: v038auth.GenesisAccounts{acc1, vaac},
	}
	supplyGenState := v036supply.GenesisState{
		Supply: supply,
	}

	migrated := v040bank.Migrate(bankGenState, authGenState, supplyGenState)
	expected := `{"params":{"send_enabled":[],"default_send_enabled":true},"balances":[{"address":"0x31ADcCDBfdf9599930Bc983877b8ac0AC990a7a5","coins":[{"denom":"stake","amount":"50"}]},{"address":"0xA328f26C3A9d1b3F4723B017b03F87EB43f3b63c","coins":[{"denom":"stake","amount":"50"}]}],"supply":[{"denom":"stake","amount":"1000"}],"denom_metadata":[]}`

	bz, err := clientCtx.Codec.MarshalJSON(migrated)
	require.NoError(t, err)
	require.Equal(t, expected, string(bz))
}
