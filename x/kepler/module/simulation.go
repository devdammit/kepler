package kepler

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/devdammit/kepler/testutil/sample"
	keplersimulation "github.com/devdammit/kepler/x/kepler/simulation"
	"github.com/devdammit/kepler/x/kepler/types"
)

// avoid unused import issue
var (
	_ = keplersimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateW3Func = "op_weight_msg_create_w_3_func"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateW3Func int = 100

	opWeightMsgW3FuncExecuted = "op_weight_msg_w_3_func_executed"
	// TODO: Determine the simulation weight value
	defaultWeightMsgW3FuncExecuted int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	keplerGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&keplerGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateW3Func int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateW3Func, &weightMsgCreateW3Func, nil,
		func(_ *rand.Rand) {
			weightMsgCreateW3Func = defaultWeightMsgCreateW3Func
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateW3Func,
		keplersimulation.SimulateMsgCreateW3Func(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgW3FuncExecuted int
	simState.AppParams.GetOrGenerate(opWeightMsgW3FuncExecuted, &weightMsgW3FuncExecuted, nil,
		func(_ *rand.Rand) {
			weightMsgW3FuncExecuted = defaultWeightMsgW3FuncExecuted
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgW3FuncExecuted,
		keplersimulation.SimulateMsgW3FuncExecuted(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateW3Func,
			defaultWeightMsgCreateW3Func,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				keplersimulation.SimulateMsgCreateW3Func(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgW3FuncExecuted,
			defaultWeightMsgW3FuncExecuted,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				keplersimulation.SimulateMsgW3FuncExecuted(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
