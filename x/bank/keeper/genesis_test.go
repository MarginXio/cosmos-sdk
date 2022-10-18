package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
)

func (suite *IntegrationTestSuite) TestExportGenesis() {
	app, ctx := suite.app, suite.ctx

	expectedMetadata := suite.getTestMetadata()
	expectedBalances, totalSupply := suite.getTestBalancesAndSupply()
	for i := range []int{1, 2} {
		app.BankKeeper.SetDenomMetaData(ctx, expectedMetadata[i])
		accAddr, err1 := sdk.AccAddressFromBech32(expectedBalances[i].Address)
		if err1 != nil {
			panic(err1)
		}
		// set balances via mint and send
		suite.
			Require().
			NoError(app.BankKeeper.MintCoins(ctx, minttypes.ModuleName, expectedBalances[i].Coins))
		suite.
			Require().
			NoError(app.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, accAddr, expectedBalances[i].Coins))
	}
	app.BankKeeper.SetParams(ctx, types.DefaultParams())

	exportGenesis := app.BankKeeper.ExportGenesis(ctx)

	suite.Require().Len(exportGenesis.Params.SendEnabled, 0)
	suite.Require().Equal(types.DefaultParams().DefaultSendEnabled, exportGenesis.Params.DefaultSendEnabled)
	suite.Require().Equal(totalSupply, exportGenesis.Supply)
	suite.Require().Equal(expectedBalances, exportGenesis.Balances)
	suite.Require().Equal(expectedMetadata, exportGenesis.DenomMetadata)
}

func (suite *IntegrationTestSuite) getTestBalancesAndSupply() ([]types.Balance, sdk.Coins) {
	addr2, _ := sdk.AccAddressFromBech32("0x494D2b9B6f0FC43B9710B26a0cE6a1c28cE0Be6f")
	addr1, _ := sdk.AccAddressFromBech32("0x5d38F92511FCa121dD5b2E6bE6D66455f3DF6737")
	addr1Balance := sdk.Coins{sdk.NewInt64Coin("testcoin3", 10)}
	addr2Balance := sdk.Coins{sdk.NewInt64Coin("testcoin1", 32), sdk.NewInt64Coin("testcoin2", 34)}

	totalSupply := addr1Balance
	totalSupply = totalSupply.Add(addr2Balance...)

	return []types.Balance{
		{Address: addr2.String(), Coins: addr2Balance},
		{Address: addr1.String(), Coins: addr1Balance},
	}, totalSupply
}

func (suite *IntegrationTestSuite) TestInitGenesis() {
	m := types.Metadata{Description: sdk.DefaultBondDenom, Base: sdk.DefaultBondDenom, Display: sdk.DefaultBondDenom}
	g := types.DefaultGenesisState()
	g.DenomMetadata = []types.Metadata{m}
	bk := suite.app.BankKeeper
	bk.InitGenesis(suite.ctx, g)

	m2, found := bk.GetDenomMetaData(suite.ctx, m.Base)
	suite.Require().True(found)
	suite.Require().Equal(m, m2)
}

func (suite *IntegrationTestSuite) TestTotalSupply() {
	// Prepare some test data.
	defaultGenesis := types.DefaultGenesisState()
	balances := []types.Balance{
		{Coins: sdk.NewCoins(sdk.NewCoin("foocoin", sdk.NewInt(1))), Address: "0x494D2b9B6f0FC43B9710B26a0cE6a1c28cE0Be6f"},
		{Coins: sdk.NewCoins(sdk.NewCoin("barcoin", sdk.NewInt(1))), Address: "0x5d38F92511FCa121dD5b2E6bE6D66455f3DF6737"},
		{Coins: sdk.NewCoins(sdk.NewCoin("foocoin", sdk.NewInt(10)), sdk.NewCoin("barcoin", sdk.NewInt(20))), Address: "0xDc6F17BBEc824FFf8F86587966B2047Db6aB7367"},
	}
	totalSupply := sdk.NewCoins(sdk.NewCoin("foocoin", sdk.NewInt(11)), sdk.NewCoin("barcoin", sdk.NewInt(21)))

	testcases := []struct {
		name        string
		genesis     *types.GenesisState
		expSupply   sdk.Coins
		expPanic    bool
		expPanicMsg string
	}{
		{
			"calculation NOT matching genesis Supply field",
			types.NewGenesisState(defaultGenesis.Params, balances, sdk.NewCoins(sdk.NewCoin("wrongcoin", sdk.NewInt(1))), defaultGenesis.DenomMetadata),
			nil, true, "genesis supply is incorrect, expected 1wrongcoin, got 21barcoin,11foocoin",
		},
		{
			"calculation matches genesis Supply field",
			types.NewGenesisState(defaultGenesis.Params, balances, totalSupply, defaultGenesis.DenomMetadata),
			totalSupply, false, "",
		},
		{
			"calculation is correct, empty genesis Supply field",
			types.NewGenesisState(defaultGenesis.Params, balances, nil, defaultGenesis.DenomMetadata),
			totalSupply, false, "",
		},
	}

	for _, tc := range testcases {
		tc := tc
		suite.Run(tc.name, func() {
			if tc.expPanic {
				suite.PanicsWithError(tc.expPanicMsg, func() { suite.app.BankKeeper.InitGenesis(suite.ctx, tc.genesis) })
			} else {
				suite.app.BankKeeper.InitGenesis(suite.ctx, tc.genesis)
				totalSupply, _, err := suite.app.BankKeeper.GetPaginatedTotalSupply(suite.ctx, &query.PageRequest{Limit: query.MaxLimit})
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expSupply, totalSupply)
			}
		})
	}
}
