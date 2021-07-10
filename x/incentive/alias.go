// Code generated by aliasgen tool (github.com/rhuairahrighairidh/aliasgen) DO NOT EDIT.

package incentive

import (
	"github.com/kava-labs/kava/x/incentive/keeper"
	"github.com/kava-labs/kava/x/incentive/types"
)

const (
	BeginningOfMonth               = keeper.BeginningOfMonth
	MidMonth                       = keeper.MidMonth
	PaymentHour                    = keeper.PaymentHour
	AttributeKeyClaimAmount        = types.AttributeKeyClaimAmount
	AttributeKeyClaimPeriod        = types.AttributeKeyClaimPeriod
	AttributeKeyClaimType          = types.AttributeKeyClaimType
	AttributeKeyClaimedBy          = types.AttributeKeyClaimedBy
	AttributeKeyRewardPeriod       = types.AttributeKeyRewardPeriod
	AttributeValueCategory         = types.AttributeValueCategory
	BondDenom                      = types.BondDenom
	DefaultParamspace              = types.DefaultParamspace
	DelegatorClaimType             = types.DelegatorClaimType
	EventTypeClaim                 = types.EventTypeClaim
	EventTypeClaimPeriod           = types.EventTypeClaimPeriod
	EventTypeClaimPeriodExpiry     = types.EventTypeClaimPeriodExpiry
	EventTypeRewardPeriod          = types.EventTypeRewardPeriod
	HardLiquidityProviderClaimType = types.HardLiquidityProviderClaimType
	Large                          = types.Large
	Medium                         = types.Medium
	ModuleName                     = types.ModuleName
	QuerierRoute                   = types.QuerierRoute
	QueryGetDelegatorRewards       = types.QueryGetDelegatorRewards
	QueryGetHardRewards            = types.QueryGetHardRewards
	QueryGetParams                 = types.QueryGetParams
	QueryGetRewardFactors          = types.QueryGetRewardFactors
	QueryGetSwapRewards            = types.QueryGetSwapRewards
	QueryGetUSDXMintingRewards     = types.QueryGetUSDXMintingRewards
	RestClaimCollateralType        = types.RestClaimCollateralType
	RestClaimOwner                 = types.RestClaimOwner
	RestClaimType                  = types.RestClaimType
	RestUnsynced                   = types.RestUnsynced
	RouterKey                      = types.RouterKey
	Small                          = types.Small
	StoreKey                       = types.StoreKey
	SwapClaimType                  = types.SwapClaimType
	USDXMintingClaimType           = types.USDXMintingClaimType
)

var (
	// function aliases
	CalculateTimeElapsed                 = keeper.CalculateTimeElapsed
	NewKeeper                            = keeper.NewKeeper
	NewQuerier                           = keeper.NewQuerier
	DefaultGenesisState                  = types.DefaultGenesisState
	DefaultParams                        = types.DefaultParams
	GetTotalVestingPeriodLength          = types.GetTotalVestingPeriodLength
	NewAccumulationTime                  = types.NewAccumulationTime
	NewAccumulator                       = types.NewAccumulator
	NewDelegatorClaim                    = types.NewDelegatorClaim
	NewGenesisRewardState                = types.NewGenesisRewardState
	NewGenesisState                      = types.NewGenesisState
	NewHardLiquidityProviderClaim        = types.NewHardLiquidityProviderClaim
	NewMsgClaimDelegatorReward           = types.NewMsgClaimDelegatorReward
	NewMsgClaimDelegatorRewardVVesting   = types.NewMsgClaimDelegatorRewardVVesting
	NewMsgClaimHardReward                = types.NewMsgClaimHardReward
	NewMsgClaimHardRewardVVesting        = types.NewMsgClaimHardRewardVVesting
	NewMsgClaimUSDXMintingReward         = types.NewMsgClaimUSDXMintingReward
	NewMsgClaimUSDXMintingRewardVVesting = types.NewMsgClaimUSDXMintingRewardVVesting
	NewMultiRewardIndex                  = types.NewMultiRewardIndex
	NewMultiRewardPeriod                 = types.NewMultiRewardPeriod
	NewMultiplier                        = types.NewMultiplier
	NewParams                            = types.NewParams
	NewPeriod                            = types.NewPeriod
	NewQueryGetRewardFactorsResponse     = types.NewQueryGetRewardFactorsResponse
	NewQueryRewardsParams                = types.NewQueryRewardsParams
	NewRewardIndex                       = types.NewRewardIndex
	NewRewardPeriod                      = types.NewRewardPeriod
	NewSwapClaim                         = types.NewSwapClaim
	NewUSDXMintingClaim                  = types.NewUSDXMintingClaim
	ParamKeyTable                        = types.ParamKeyTable
	RegisterCodec                        = types.RegisterCodec

	// variable aliases
	DefaultActive                                 = types.DefaultActive
	DefaultClaimEnd                               = types.DefaultClaimEnd
	DefaultDelegatorClaims                        = types.DefaultDelegatorClaims
	DefaultGenesisRewardState                     = types.DefaultGenesisRewardState
	DefaultHardClaims                             = types.DefaultHardClaims
	DefaultMultiRewardPeriods                     = types.DefaultMultiRewardPeriods
	DefaultMultipliers                            = types.DefaultMultipliers
	DefaultRewardPeriods                          = types.DefaultRewardPeriods
	DefaultSwapClaims                             = types.DefaultSwapClaims
	DefaultUSDXClaims                             = types.DefaultUSDXClaims
	DelegatorClaimKeyPrefix                       = types.DelegatorClaimKeyPrefix
	DelegatorRewardIndexesKeyPrefix               = types.DelegatorRewardIndexesKeyPrefix
	ErrAccountNotFound                            = types.ErrAccountNotFound
	ErrClaimExpired                               = types.ErrClaimExpired
	ErrClaimNotFound                              = types.ErrClaimNotFound
	ErrDecreasingRewardFactor                     = types.ErrDecreasingRewardFactor
	ErrInsufficientModAccountBalance              = types.ErrInsufficientModAccountBalance
	ErrInvalidAccountType                         = types.ErrInvalidAccountType
	ErrInvalidClaimOwner                          = types.ErrInvalidClaimOwner
	ErrInvalidClaimType                           = types.ErrInvalidClaimType
	ErrInvalidMultiplier                          = types.ErrInvalidMultiplier
	ErrNoClaimsFound                              = types.ErrNoClaimsFound
	ErrRewardPeriodNotFound                       = types.ErrRewardPeriodNotFound
	ErrZeroClaim                                  = types.ErrZeroClaim
	GovDenom                                      = types.GovDenom
	HardBorrowRewardIndexesKeyPrefix              = types.HardBorrowRewardIndexesKeyPrefix
	HardLiquidityClaimKeyPrefix                   = types.HardLiquidityClaimKeyPrefix
	HardLiquidityRewardDenom                      = types.HardLiquidityRewardDenom
	HardSupplyRewardIndexesKeyPrefix              = types.HardSupplyRewardIndexesKeyPrefix
	IncentiveMacc                                 = types.IncentiveMacc
	KeyClaimEnd                                   = types.KeyClaimEnd
	KeyDelegatorRewardPeriods                     = types.KeyDelegatorRewardPeriods
	KeyHardBorrowRewardPeriods                    = types.KeyHardBorrowRewardPeriods
	KeyHardSupplyRewardPeriods                    = types.KeyHardSupplyRewardPeriods
	KeyMultipliers                                = types.KeyMultipliers
	KeySwapRewardPeriods                          = types.KeySwapRewardPeriods
	KeyUSDXMintingRewardPeriods                   = types.KeyUSDXMintingRewardPeriods
	ModuleCdc                                     = types.ModuleCdc
	PreviousDelegatorRewardAccrualTimeKeyPrefix   = types.PreviousDelegatorRewardAccrualTimeKeyPrefix
	PreviousHardBorrowRewardAccrualTimeKeyPrefix  = types.PreviousHardBorrowRewardAccrualTimeKeyPrefix
	PreviousHardSupplyRewardAccrualTimeKeyPrefix  = types.PreviousHardSupplyRewardAccrualTimeKeyPrefix
	PreviousSwapRewardAccrualTimeKeyPrefix        = types.PreviousSwapRewardAccrualTimeKeyPrefix
	PreviousUSDXMintingRewardAccrualTimeKeyPrefix = types.PreviousUSDXMintingRewardAccrualTimeKeyPrefix
	PrincipalDenom                                = types.PrincipalDenom
	SwapClaimKeyPrefix                            = types.SwapClaimKeyPrefix
	SwapRewardIndexesKeyPrefix                    = types.SwapRewardIndexesKeyPrefix
	USDXMintingClaimKeyPrefix                     = types.USDXMintingClaimKeyPrefix
	USDXMintingRewardDenom                        = types.USDXMintingRewardDenom
	USDXMintingRewardFactorKeyPrefix              = types.USDXMintingRewardFactorKeyPrefix
)

type (
	Hooks                             = keeper.Hooks
	Keeper                            = keeper.Keeper
	AccountKeeper                     = types.AccountKeeper
	AccumulationTime                  = types.AccumulationTime
	AccumulationTimes                 = types.AccumulationTimes
	Accumulator                       = types.Accumulator
	BaseClaim                         = types.BaseClaim
	BaseMultiClaim                    = types.BaseMultiClaim
	CDPHooks                          = types.CDPHooks
	CdpKeeper                         = types.CdpKeeper
	Claim                             = types.Claim
	Claims                            = types.Claims
	DelegatorClaim                    = types.DelegatorClaim
	DelegatorClaims                   = types.DelegatorClaims
	GenesisRewardState                = types.GenesisRewardState
	GenesisState                      = types.GenesisState
	HARDHooks                         = types.HARDHooks
	HardKeeper                        = types.HardKeeper
	HardLiquidityProviderClaim        = types.HardLiquidityProviderClaim
	HardLiquidityProviderClaims       = types.HardLiquidityProviderClaims
	MsgClaimDelegatorReward           = types.MsgClaimDelegatorReward
	MsgClaimDelegatorRewardVVesting   = types.MsgClaimDelegatorRewardVVesting
	MsgClaimHardReward                = types.MsgClaimHardReward
	MsgClaimHardRewardVVesting        = types.MsgClaimHardRewardVVesting
	MsgClaimUSDXMintingReward         = types.MsgClaimUSDXMintingReward
	MsgClaimUSDXMintingRewardVVesting = types.MsgClaimUSDXMintingRewardVVesting
	MultiRewardIndex                  = types.MultiRewardIndex
	MultiRewardIndexes                = types.MultiRewardIndexes
	MultiRewardPeriod                 = types.MultiRewardPeriod
	MultiRewardPeriods                = types.MultiRewardPeriods
	Multiplier                        = types.Multiplier
	MultiplierName                    = types.MultiplierName
	Multipliers                       = types.Multipliers
	ParamSubspace                     = types.ParamSubspace
	Params                            = types.Params
	QueryGetRewardFactorsResponse     = types.QueryGetRewardFactorsResponse
	QueryRewardsParams                = types.QueryRewardsParams
	RewardIndex                       = types.RewardIndex
	RewardIndexes                     = types.RewardIndexes
	RewardPeriod                      = types.RewardPeriod
	RewardPeriods                     = types.RewardPeriods
	StakingKeeper                     = types.StakingKeeper
	SupplyKeeper                      = types.SupplyKeeper
	SwapClaim                         = types.SwapClaim
	SwapClaims                        = types.SwapClaims
	SwapKeeper                        = types.SwapKeeper
	USDXMintingClaim                  = types.USDXMintingClaim
	USDXMintingClaims                 = types.USDXMintingClaims
)
