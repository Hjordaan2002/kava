package keeper_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/kava-labs/kava/app"
	"github.com/kava-labs/kava/x/incentive/keeper"
	"github.com/kava-labs/kava/x/incentive/types"
)

// Test suite used for all keeper tests
type KeeperTestSuite struct {
	suite.Suite

	keeper keeper.Keeper

	app app.TestApp
	ctx sdk.Context

	genesisTime time.Time
	addrs       []sdk.AccAddress
}

// SetupTest is run automatically before each suite test
func (suite *KeeperTestSuite) SetupTest() {
	config := sdk.GetConfig()
	app.SetBech32AddressPrefixes(config)

	_, suite.addrs = app.GeneratePrivKeyAddressPairs(5)

	suite.genesisTime = time.Date(2020, 12, 15, 14, 0, 0, 0, time.UTC)
}

func (suite *KeeperTestSuite) SetupApp() {
	suite.app = app.NewTestApp()

	suite.keeper = suite.app.GetIncentiveKeeper()

	suite.ctx = suite.app.NewContext(true, abci.Header{Height: 1, Time: suite.genesisTime})
}

func (suite *KeeperTestSuite) TestGetSetDeleteUSDXMintingClaim() {
	suite.SetupApp()
	c := types.NewUSDXMintingClaim(suite.addrs[0], c("ukava", 1000000), types.RewardIndexes{types.NewRewardIndex("bnb-a", sdk.ZeroDec())})
	_, found := suite.keeper.GetUSDXMintingClaim(suite.ctx, c.Owner)
	suite.Require().False(found)
	suite.Require().NotPanics(func() {
		suite.keeper.SetUSDXMintingClaim(suite.ctx, c)
	})
	testC, found := suite.keeper.GetUSDXMintingClaim(suite.ctx, c.Owner)
	suite.Require().True(found)
	suite.Require().Equal(c, testC)
	suite.Require().NotPanics(func() {
		suite.keeper.DeleteUSDXMintingClaim(suite.ctx, c.Owner)
	})
	_, found = suite.keeper.GetUSDXMintingClaim(suite.ctx, c.Owner)
	suite.Require().False(found)
}

// var nonEmptyRewardIndexes = types.RewardIndexes{
// 	types.NewRewardIndex("hard", d("0.1")),
// 	types.NewRewardIndex("ukava", d("0.2")),
// }
// var singleRewardIndexes = types.RewardIndexes{
// 	types.NewRewardIndex("hard", d("0.1")),
// }
// var nonEmptyMultiRewardIndexes = types.MultiRewardIndexes{
// 	// sorted by denom
// 	types.NewMultiRewardIndex("bnb", singleRewardIndexes), // use single to ensure indexes are different
// 	types.NewMultiRewardIndex("btcb", nonEmptyRewardIndexes),
// }

func (suite *KeeperTestSuite) TestGetSetDeleteHardLiquidityProviderClaim() {
	suite.SetupApp()

	claim := types.NewHardLiquidityProviderClaim(
		suite.addrs[0],
		cs(c("ukava", 1e9), c("hard", 1e9)),
		nonEmptyMultiRewardIndexes,
		nonEmptyMultiRewardIndexes,
		nonEmptyRewardIndexes,
	)
	_, found := suite.keeper.GetHardLiquidityProviderClaim(suite.ctx, claim.Owner)
	suite.Require().False(found)
	suite.Require().NotPanics(func() {
		suite.keeper.SetHardLiquidityProviderClaim(suite.ctx, claim)
	})
	storedClaim, found := suite.keeper.GetHardLiquidityProviderClaim(suite.ctx, claim.Owner)
	suite.Require().True(found)
	suite.Require().Equal(claim, storedClaim)
	suite.Require().NotPanics(func() {
		suite.keeper.DeleteHardLiquidityProviderClaim(suite.ctx, claim.Owner)
	})
	_, found = suite.keeper.GetHardLiquidityProviderClaim(suite.ctx, claim.Owner)
	suite.Require().False(found)
}

func (suite *KeeperTestSuite) TestIterateUSDXMintingClaims() {
	suite.SetupApp()
	for i := 0; i < len(suite.addrs); i++ {
		c := types.NewUSDXMintingClaim(suite.addrs[i], c("ukava", 100000), types.RewardIndexes{types.NewRewardIndex("bnb-a", sdk.ZeroDec())})
		suite.Require().NotPanics(func() {
			suite.keeper.SetUSDXMintingClaim(suite.ctx, c)
		})
	}
	claims := types.USDXMintingClaims{}
	suite.keeper.IterateUSDXMintingClaims(suite.ctx, func(c types.USDXMintingClaim) bool {
		claims = append(claims, c)
		return false
	})
	suite.Require().Equal(len(suite.addrs), len(claims))

	claims = suite.keeper.GetAllUSDXMintingClaims(suite.ctx)
	suite.Require().Equal(len(suite.addrs), len(claims))
}

type accrualtime struct {
	denom string
	time  time.Time
}

var nonEmptyAccrualTimes = []accrualtime{
	{
		denom: "",
		time:  time.Time{},
	},
	{
		denom: "btcb",
		time:  time.Date(1998, 1, 1, 0, 0, 0, 1, time.UTC),
	},
}

func (suite *KeeperTestSuite) TestIterateUSDXMintingAccrualTimes_Empty() {
	suite.SetupApp()

	var expectedAccrualTimes []accrualtime

	var actualAccrualTimes []accrualtime
	suite.keeper.IterateUSDXMintingAccrualTimes(suite.ctx, func(denom string, accrualTime time.Time) bool {
		actualAccrualTimes = append(actualAccrualTimes, accrualtime{denom: denom, time: accrualTime})
		return false
	})

	suite.Equal(expectedAccrualTimes, actualAccrualTimes)
}

func (suite *KeeperTestSuite) TestIterateUSDXMintingAccrualTimes_NotEmpty() {
	suite.SetupApp()

	expectedAccrualTimes := nonEmptyAccrualTimes

	for _, at := range expectedAccrualTimes {
		suite.keeper.SetPreviousUSDXMintingAccrualTime(suite.ctx, at.denom, at.time)
	}

	var actualAccrualTimes []accrualtime
	suite.keeper.IterateUSDXMintingAccrualTimes(suite.ctx, func(denom string, accrualTime time.Time) bool {
		actualAccrualTimes = append(actualAccrualTimes, accrualtime{denom: denom, time: accrualTime})
		return false
	})

	suite.Equal(expectedAccrualTimes, actualAccrualTimes)
}

func (suite *KeeperTestSuite) TestIterateUSDXMintingRewardFactors_NotEmpty() {
	suite.SetupApp()

	expectedRewardIndexes := nonEmptyRewardIndexes
	for _, rf := range expectedRewardIndexes {
		suite.keeper.SetUSDXMintingRewardFactor(suite.ctx, rf.CollateralType, rf.RewardFactor)
	}

	var actualRewardIndexes types.RewardIndexes
	suite.keeper.IterateUSDXMintingRewardFactors(suite.ctx, func(denom string, factor sdk.Dec) bool {
		actualRewardIndexes = append(actualRewardIndexes, types.NewRewardIndex(denom, factor))
		return false
	})

	suite.Equal(expectedRewardIndexes, actualRewardIndexes)
}

func (suite *KeeperTestSuite) TestIterateHardDelegatorRewardAccrualTimes_NotEmpty() {
	suite.SetupApp()

	expectedAccrualTimes := nonEmptyAccrualTimes

	for _, at := range expectedAccrualTimes {
		suite.keeper.SetPreviousHardDelegatorRewardAccrualTime(suite.ctx, at.denom, at.time)
	}

	var actualAccrualTimes []accrualtime
	suite.keeper.IterateHardDelegatorRewardAccrualTimes(suite.ctx, func(denom string, accrualTime time.Time) bool {
		actualAccrualTimes = append(actualAccrualTimes, accrualtime{denom: denom, time: accrualTime})
		return false
	})

	suite.Equal(expectedAccrualTimes, actualAccrualTimes)
}

func (suite *KeeperTestSuite) TestIterateHardDelegatorRewardFactors_NotEmpty() {
	suite.SetupApp()

	expectedRewardIndexes := nonEmptyRewardIndexes
	for _, rf := range expectedRewardIndexes {
		suite.keeper.SetHardDelegatorRewardFactor(suite.ctx, rf.CollateralType, rf.RewardFactor)
	}

	var actualRewardIndexes types.RewardIndexes
	suite.keeper.IterateHardDelegatorRewardFactors(suite.ctx, func(denom string, factor sdk.Dec) bool {
		actualRewardIndexes = append(actualRewardIndexes, types.NewRewardIndex(denom, factor))
		return false
	})

	suite.Equal(expectedRewardIndexes, actualRewardIndexes)
}

func (suite *KeeperTestSuite) TestIterateHardSupplyRewardAccrualTimes_NotEmpty() {
	suite.SetupApp()

	expectedAccrualTimes := nonEmptyAccrualTimes

	for _, at := range expectedAccrualTimes {
		suite.keeper.SetPreviousHardSupplyRewardAccrualTime(suite.ctx, at.denom, at.time)
	}

	var actualAccrualTimes []accrualtime
	suite.keeper.IterateHardSupplyRewardAccrualTimes(suite.ctx, func(denom string, accrualTime time.Time) bool {
		actualAccrualTimes = append(actualAccrualTimes, accrualtime{denom: denom, time: accrualTime})
		return false
	})

	suite.Equal(expectedAccrualTimes, actualAccrualTimes)
}

func (suite *KeeperTestSuite) TestIterateHardSupplyRewardFactors_NotEmpty() {
	expectedRewardIndexes := nonEmptyMultiRewardIndexes
	for _, rf := range expectedRewardIndexes {
		suite.keeper.SetHardSupplyRewardIndexes(suite.ctx, rf.CollateralType, rf.RewardIndexes)
	}

	var actualRewardIndexes types.MultiRewardIndexes
	suite.keeper.IterateHardSupplyRewardIndexes(suite.ctx, func(denom string, indexes types.RewardIndexes) bool {
		actualRewardIndexes = append(actualRewardIndexes, types.NewMultiRewardIndex(denom, indexes))
		return false
	})

	suite.Equal(expectedRewardIndexes, actualRewardIndexes)
}

func (suite *KeeperTestSuite) TestIterateHardBorrowrRewardAccrualTimes_NotEmpty() {
	suite.SetupApp()

	expectedAccrualTimes := nonEmptyAccrualTimes

	for _, at := range expectedAccrualTimes {
		suite.keeper.SetPreviousHardBorrowRewardAccrualTime(suite.ctx, at.denom, at.time)
	}

	var actualAccrualTimes []accrualtime
	suite.keeper.IterateHardBorrowRewardAccrualTimes(suite.ctx, func(denom string, accrualTime time.Time) bool {
		actualAccrualTimes = append(actualAccrualTimes, accrualtime{denom: denom, time: accrualTime})
		return false
	})

	suite.Equal(expectedAccrualTimes, actualAccrualTimes)
}

func (suite *KeeperTestSuite) TestIterateHardBorrowRewardFactors_NotEmpty() {
	suite.SetupApp()

	expectedRewardIndexes := nonEmptyMultiRewardIndexes
	for _, rf := range expectedRewardIndexes {
		suite.keeper.SetHardBorrowRewardIndexes(suite.ctx, rf.CollateralType, rf.RewardIndexes)
	}

	var actualRewardIndexes types.MultiRewardIndexes
	suite.keeper.IterateHardBorrowRewardIndexes(suite.ctx, func(denom string, indexes types.RewardIndexes) bool {
		actualRewardIndexes = append(actualRewardIndexes, types.NewMultiRewardIndex(denom, indexes))
		return false
	})

	suite.Equal(expectedRewardIndexes, actualRewardIndexes)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
