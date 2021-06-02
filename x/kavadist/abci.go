package kavadist

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var v0142UpgradeTime = time.Date(2021, 6, 10, 14, 0, 0, 0, time.UTC)

func BeginBlocker(ctx sdk.Context, k Keeper) {
	err := k.MintPeriodInflation(ctx)
	if err != nil {
		panic(err)
	}
	if ctx.BlockTime().After(v0142UpgradeTime) {
		RegisterMultiSpend(ModuleCdc)
	}
}
