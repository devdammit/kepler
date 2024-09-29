package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/devdammit/kepler/x/kepler/keeper"
	"github.com/devdammit/kepler/x/kepler/types"
)

func SimulateMsgCreateW3Func(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCreateW3Func{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the CreateW3Func simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "CreateW3Func simulation not implemented"), nil, nil
	}
}
