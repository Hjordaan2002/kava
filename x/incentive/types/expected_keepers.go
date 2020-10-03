package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authexported "github.com/cosmos/cosmos-sdk/x/auth/exported"
	supplyexported "github.com/cosmos/cosmos-sdk/x/supply/exported"

	cdptypes "github.com/kava-labs/kava/x/cdp/types"
)

// SupplyKeeper defines the expected supply keeper for module accounts
type SupplyKeeper interface {
	GetModuleAccount(ctx sdk.Context, name string) supplyexported.ModuleAccountI

	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
}

// CdpKeeper defines the expected cdp keeper for interacting with cdps
type CdpKeeper interface {
	IterateCdpsByCollateralType(ctx sdk.Context, collateralType string, cb func(cdp cdptypes.CDP) (stop bool))
	GetTotalPrincipal(ctx sdk.Context, collateralType string, principalDenom string) (total sdk.Int)
}

// AccountKeeper defines the expected keeper interface for interacting with account
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authexported.Account
	SetAccount(ctx sdk.Context, acc authexported.Account)
	GetNextAccountNumber(ctx sdk.Context) uint64
}
