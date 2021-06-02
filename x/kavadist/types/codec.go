package types

import (
	"bytes"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
)

// ModuleCdc generic sealed codec to be used throughout module
var ModuleCdc *codec.Codec

func init() {
	cdc := codec.New()
	RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	ModuleCdc = cdc
}

// RegisterCodec registers the necessary types for cdp module
func RegisterCodec(cdc *codec.Codec) {
	RegisterMultiSpend(cdc)
}

func RegisterMultiSpend(cdc *codec.Codec) {
	buf := new(bytes.Buffer)
	cdc.PrintTypes(buf)
	if strings.Contains(buf.String(), "kava/CommunityPoolMultiSpendProposal") {
		return
	}
	cdc.RegisterConcrete(CommunityPoolMultiSpendProposal{}, "kava/CommunityPoolMultiSpendProposal", nil)
}
