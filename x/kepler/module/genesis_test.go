package kepler_test

import (
	"testing"

	keepertest "github.com/devdammit/kepler/testutil/keeper"
	"github.com/devdammit/kepler/testutil/nullify"
	kepler "github.com/devdammit/kepler/x/kepler/module"
	"github.com/devdammit/kepler/x/kepler/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.KeplerKeeper(t)
	kepler.InitGenesis(ctx, k, genesisState)
	got := kepler.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
