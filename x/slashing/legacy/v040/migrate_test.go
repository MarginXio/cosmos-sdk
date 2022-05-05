package v040_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	v039slashing "github.com/cosmos/cosmos-sdk/x/slashing/legacy/v039"
	v040slashing "github.com/cosmos/cosmos-sdk/x/slashing/legacy/v040"
)

func TestMigrate(t *testing.T) {
	encodingConfig := simapp.MakeTestEncodingConfig()
	clientCtx := client.Context{}.
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithJSONCodec(encodingConfig.Marshaler)

	addr1, err := sdk.ConsAddressFromBech32("0x7D712D9Ac341c87fEC61551B97Ab35125EF17f35")
	require.NoError(t, err)
	addr2, err := sdk.ConsAddressFromBech32("0x7E6B8a0740b3e04C10bCD72b15fb10dfcC5493AD")
	require.NoError(t, err)

	gs := v039slashing.GenesisState{
		Params: v039slashing.DefaultParams(),
		SigningInfos: map[string]v039slashing.ValidatorSigningInfo{
			"0x7E6B8a0740b3e04C10bCD72b15fb10dfcC5493AD": {
				Address:             addr2,
				IndexOffset:         615501,
				MissedBlocksCounter: 1,
				Tombstoned:          false,
			},
			"0x7D712D9Ac341c87fEC61551B97Ab35125EF17f35": {
				Address:             addr1,
				IndexOffset:         2,
				MissedBlocksCounter: 2,
				Tombstoned:          false,
			},
		},
		MissedBlocks: map[string][]v039slashing.MissedBlock{
			"0x7E6B8a0740b3e04C10bCD72b15fb10dfcC5493AD": {
				{
					Index:  2,
					Missed: true,
				},
			},
			"0x7D712D9Ac341c87fEC61551B97Ab35125EF17f35": {
				{
					Index:  3,
					Missed: true,
				},
				{
					Index:  4,
					Missed: true,
				},
			},
		},
	}

	migrated := v040slashing.Migrate(gs)
	// Check that in `signing_infos` and `missed_blocks`, the address
	// 0x7D712D9Ac341c87fEC61551B97Ab35125EF17f35
	// should always come before the address
	// 0x7E6B8a0740b3e04C10bCD72b15fb10dfcC5493AD
	// (in alphabetic order, basically).
	expected := `{
  "missed_blocks": [
    {
      "address": "0x7D712D9Ac341c87fEC61551B97Ab35125EF17f35",
      "missed_blocks": [
        {
          "index": "3",
          "missed": true
        },
        {
          "index": "4",
          "missed": true
        }
      ]
    },
    {
      "address": "0x7E6B8a0740b3e04C10bCD72b15fb10dfcC5493AD",
      "missed_blocks": [
        {
          "index": "2",
          "missed": true
        }
      ]
    }
  ],
  "params": {
    "downtime_jail_duration": "600s",
    "min_signed_per_window": "0.500000",
    "signed_blocks_window": "100",
    "slash_fraction_double_sign": "0.050000",
    "slash_fraction_downtime": "0.010000"
  },
  "signing_infos": [
    {
      "address": "0x7D712D9Ac341c87fEC61551B97Ab35125EF17f35",
      "validator_signing_info": {
        "address": "0x7D712D9Ac341c87fEC61551B97Ab35125EF17f35",
        "index_offset": "2",
        "jailed_until": "0001-01-01T00:00:00Z",
        "missed_blocks_counter": "2",
        "start_height": "0",
        "tombstoned": false
      }
    },
    {
      "address": "0x7E6B8a0740b3e04C10bCD72b15fb10dfcC5493AD",
      "validator_signing_info": {
        "address": "0x7E6B8a0740b3e04C10bCD72b15fb10dfcC5493AD",
        "index_offset": "615501",
        "jailed_until": "0001-01-01T00:00:00Z",
        "missed_blocks_counter": "1",
        "start_height": "0",
        "tombstoned": false
      }
    }
  ]
}`

	bz, err := clientCtx.Codec.MarshalJSON(migrated)
	require.NoError(t, err)

	// Indent the JSON bz correctly.
	var jsonObj map[string]interface{}
	err = json.Unmarshal(bz, &jsonObj)
	require.NoError(t, err)
	indentedBz, err := json.MarshalIndent(jsonObj, "", "  ")
	require.NoError(t, err)

	require.Equal(t, expected, string(indentedBz))
}
