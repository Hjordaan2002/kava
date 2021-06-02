package kavadist

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/kava-labs/kava/x/kavadist/keeper"
	"github.com/kava-labs/kava/x/kavadist/types"
)

var v0142UpgradeTime = time.Date(2021, 6, 10, 14, 0, 0, 0, time.UTC)

// NewHandler creates an sdk.Handler for kavadist messages
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s message type: %T", ModuleName, msg)
		}
	}
}

func NewCommunityPoolMultiSpendProposalHandler(k Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case types.CommunityPoolMultiSpendProposal:
			if ctx.BlockTime().After(v0142UpgradeTime) {
				return keeper.HandleCommunityPoolMultiSpendProposal(ctx, k, c)
			}
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized kavadist proposal content type: %T", c)
		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized kavadist proposal content type: %T", c)
		}
	}
}
