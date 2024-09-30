package kepler

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/devdammit/kepler/api/kepler/kepler"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "ListW3Func",
					Use:       "list-w3func",
					Short:     "Query list-w3func",
				},

				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateW3Func",
					Use:            "create-w3func [cond] [code]",
					Short:          "Send a create-w3func tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "cond"}, {ProtoField: "code"}},
				},
				{
					RpcMethod:      "W3FuncExecuted",
					Use:            "w-3-func-executed [id] [executed-at]",
					Short:          "Send a w3func-executed tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "executedAt"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
