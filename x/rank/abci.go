package rank

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/deep-foundation/deep-chain/x/rank/keeper"
)

func EndBlocker(ctx sdk.Context, k *keeper.StateKeeper) {
	k.EndBlocker(ctx)
}
