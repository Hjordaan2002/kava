package simulation

import (
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authexported "github.com/cosmos/cosmos-sdk/x/auth/exported"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	appparams "github.com/kava-labs/kava/app/params"
	"github.com/kava-labs/kava/x/swap/keeper"
	"github.com/kava-labs/kava/x/swap/types"
)

var (
	//nolint
	noOpMsg             = simulation.NoOpMsg(types.ModuleName)
	errorNotEnoughCoins = errors.New("account doesn't have enough coins")
)

// Simulation operation weights constants
const (
	OpWeightMsgDeposit = "op_weight_msg_deposit"
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(
	appParams simulation.AppParams, cdc *codec.Codec, ak types.AccountKeeper, k keeper.Keeper,
) simulation.WeightedOperations {
	var weightMsgDeposit int

	appParams.GetOrGenerate(cdc, OpWeightMsgDeposit, &weightMsgDeposit, nil,
		func(_ *rand.Rand) {
			weightMsgDeposit = appparams.DefaultWeightMsgDeposit
		},
	)

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightMsgDeposit,
			SimulateMsgDeposit(ak, k),
		),
	}
}

// SimulateMsgDeposit generates a MsgDeposit
func SimulateMsgDeposit(ak types.AccountKeeper, k keeper.Keeper) simulation.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simulation.Account, chainID string,
	) (simulation.OperationMsg, []simulation.FutureOperation, error) {
		// Get possible pools and shuffle so that deposits are evenly distributed across pools
		params := k.GetParams(ctx)
		allowedPools := params.AllowedPools
		r.Shuffle(len(allowedPools), func(i, j int) {
			allowedPools[i], allowedPools[j] = allowedPools[j], allowedPools[i]
		})

		// Find an account-pool pair that is likely to result in a successful deposit
		blockTime := ctx.BlockHeader().Time
		depositor, allowedPool, found := findValidAccountAllowedPoolPair(accs, allowedPools, func(acc simulation.Account, pool types.AllowedPool) bool {
			account := ak.GetAccount(ctx, acc.Address)

			err := validateDepositor(ctx, k, pool, account, blockTime)
			if err == errorNotEnoughCoins {
				return false // keep searching
			} else if err != nil {
				panic(err) // raise errors
			}
			return true // found valid pair
		})
		if !found {
			return simulation.NewOperationMsgBasic(types.ModuleName, "no-operation (no valid allowed pool and depositor)", "", false, nil), nil, nil
		}

		// Construct initial msg using slippage, deadline, and depositor's address
		slippageRaw, err := RandIntInclusive(r, sdk.OneInt(), sdk.NewInt(99))
		if err != nil {
			panic(err)
		}
		slippage := slippageRaw.ToDec().Quo(sdk.NewDec(100))

		durationNanoseconds, err := RandIntInclusive(r,
			sdk.NewInt((time.Second * 10).Nanoseconds()), // ten seconds
			sdk.NewInt((time.Hour * 24).Nanoseconds()),   // one day
		)
		if err != nil {
			panic(err)
		}
		extraTime := time.Duration(durationNanoseconds.Int64())
		deadline := blockTime.Add(extraTime).Unix()

		depositorAcc := ak.GetAccount(ctx, depositor.Address)
		depositorCoins := depositorAcc.SpendableCoins(blockTime)

		msg := types.NewMsgDeposit(depositorAcc.GetAddress(), sdk.Coin{}, sdk.Coin{}, slippage, deadline)

		// Populate msg with randomized token amounts
		pool, found := k.GetPool(ctx, allowedPool.Name())
		if !found { // Pool doesn't exist: first deposit
			depositTokenA := randCoinFromCoins(r, depositorCoins, allowedPool.TokenA)
			msg.TokenA = depositTokenA

			depositTokenB := randCoinFromCoins(r, depositorCoins, allowedPool.TokenB)
			msg.TokenB = depositTokenB
		} else { // Pool exists: successive deposit
			var denomX string // Denom X is the token denom in the pool with the larger amount
			var denomY string //  Denom Y is the token denom in the pool with the larger amount
			if pool.ReservesA.Amount.GTE(pool.ReservesB.Amount) {
				denomX = pool.ReservesA.Denom
				denomY = pool.ReservesB.Denom
			} else {
				denomX = pool.ReservesB.Denom
				denomY = pool.ReservesA.Denom
			}
			depositTokenY := randCoinFromCoins(r, depositorCoins, denomY)
			msg.TokenA = depositTokenY

			// Calculate the pool's slippage ratio and use it to build other coin
			ratio := pool.Reserves().AmountOf(denomX).ToDec().Quo(pool.Reserves().AmountOf(denomY).ToDec())
			amtTokenX := depositTokenY.Amount.ToDec().Mul(ratio).RoundInt()
			depositTokenX := sdk.NewCoin(denomX, amtTokenX)
			if depositorCoins.AmountOf(denomX).LT(amtTokenX) {
				return simulation.NewOperationMsgBasic(types.ModuleName, "no-operation (depositor has insufficient coins)", "", false, nil), nil, nil
			}
			msg.TokenB = depositTokenX
		}

		err = msg.ValidateBasic()
		if err != nil {
			return noOpMsg, nil, nil
		}

		tx := helpers.GenTx(
			[]sdk.Msg{msg},
			sdk.NewCoins(),
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{depositorAcc.GetAccountNumber()},
			[]uint64{depositorAcc.GetSequence()},
			depositor.PrivKey,
		)

		_, result, err := app.Deliver(tx)
		if err != nil {
			// to aid debugging, add the stack trace to the comment field of the returned opMsg
			return simulation.NewOperationMsg(msg, false, fmt.Sprintf("%+v", err)), nil, err
		}
		return simulation.NewOperationMsg(msg, true, result.Log), nil, nil
	}
}

// From a set of coins return a coin of the specified denom with 1-10% of the total amount
func randCoinFromCoins(r *rand.Rand, coins sdk.Coins, denom string) sdk.Coin {
	percentOfBalance, err := RandIntInclusive(r, sdk.OneInt(), sdk.NewInt(10))
	if err != nil {
		panic(err)
	}
	balance := coins.AmountOf(denom)
	amtToken := balance.Mul(percentOfBalance).Quo(sdk.NewInt(100))
	return sdk.NewCoin(denom, amtToken)
}

func validateDepositor(ctx sdk.Context, k keeper.Keeper, allowedPool types.AllowedPool,
	depositor authexported.Account, blockTime time.Time) error {
	depositorCoins := depositor.SpendableCoins(blockTime)
	tokenABalance := depositorCoins.AmountOf(allowedPool.TokenA)
	tokenBBalance := depositorCoins.AmountOf(allowedPool.TokenB)

	_, found := k.GetPool(ctx, allowedPool.Name())
	if !found {
		if tokenABalance.IsZero() || tokenBBalance.IsZero() {
			return errorNotEnoughCoins
		}
	} else {
		if tokenABalance.IsZero() && tokenBBalance.IsZero() {
			return errorNotEnoughCoins
		}
	}

	return nil
}

// findValidAccountAllowedPoolPair finds an account for which the callback func returns true
func findValidAccountAllowedPoolPair(accounts []simulation.Account, pools types.AllowedPools,
	cb func(simulation.Account, types.AllowedPool) bool) (simulation.Account, types.AllowedPool, bool) {
	for _, pool := range pools {
		for _, acc := range accounts {
			if isValid := cb(acc, pool); isValid {
				return acc, pool, true
			}
		}
	}
	return simulation.Account{}, types.AllowedPool{}, false
}

// RandIntInclusive randomly generates an sdk.Int in the range [inclusiveMin, inclusiveMax]. It works for negative and positive integers.
func RandIntInclusive(r *rand.Rand, inclusiveMin, inclusiveMax sdk.Int) (sdk.Int, error) {
	if inclusiveMin.GT(inclusiveMax) {
		return sdk.Int{}, fmt.Errorf("min larger than max")
	}
	return RandInt(r, inclusiveMin, inclusiveMax.Add(sdk.OneInt()))
}

// RandInt randomly generates an sdk.Int in the range [inclusiveMin, exclusiveMax). It works for negative and positive integers.
func RandInt(r *rand.Rand, inclusiveMin, exclusiveMax sdk.Int) (sdk.Int, error) {
	// validate input
	if inclusiveMin.GTE(exclusiveMax) {
		return sdk.Int{}, fmt.Errorf("min larger or equal to max")
	}
	// shift the range to start at 0
	shiftedRange := exclusiveMax.Sub(inclusiveMin) // should always be positive given the check above
	// randomly pick from the shifted range
	shiftedRandInt := sdk.NewIntFromBigInt(new(big.Int).Rand(r, shiftedRange.BigInt()))
	// shift back to the original range
	return shiftedRandInt.Add(inclusiveMin), nil
}
